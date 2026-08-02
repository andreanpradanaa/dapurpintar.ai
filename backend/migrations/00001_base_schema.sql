-- +goose Up
-- M5-001 base schema (approved). Creates all aggregate tables in dependency
-- order per docs/database/m5-migrations.md. Owned bounded context is noted per
-- table. auth_sessions is added by DP-FEAT-001 to satisfy M4-DEC-003: durable
-- session and revocation authority in PostgreSQL.
--
-- UUID primary keys use gen_random_uuid() from core (PostgreSQL 13+).
-- All instants are timestamptz stored in UTC (M4-DEC-007). Text lifecycle
-- columns use check constraints for MVP migration friction.

-- Identity and Access
create table accounts (
    id                uuid primary key default gen_random_uuid(),
    email             text not null unique,
    password_hash     text not null,
    status            text not null default 'pending'
                      check (status in ('pending', 'active', 'restricted', 'closed')),
    timezone          text not null default 'Asia/Jakarta',
    email_verified_at timestamptz,
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    deleted_at        timestamptz
);

create index accounts_status_idx on accounts (status) where deleted_at is null;

-- Identity and Access: durable refresh-session authority (M4-DEC-003).
-- Stores only a protected representation of the refresh secret and session
-- metadata required for revocation, expiry, rotation, and reuse detection.
create table auth_sessions (
    id                  uuid primary key default gen_random_uuid(),
    account_id          uuid not null references accounts (id),
    refresh_secret_hash text not null,
    family_id           uuid not null,
    user_agent_hash     text,
    ip_hash             text,
    expires_at          timestamptz not null,
    revoked_at          timestamptz,
    replaced_by         uuid,
    created_at          timestamptz not null default now(),
    updated_at          timestamptz not null default now()
);

create index auth_sessions_account_idx on auth_sessions (account_id);
create index auth_sessions_secret_idx on auth_sessions (refresh_secret_hash);
create index auth_sessions_expiry_idx on auth_sessions (expires_at);
create index auth_sessions_family_idx on auth_sessions (family_id);

-- User Context and Preferences
create table user_profiles (
    id           uuid primary key default gen_random_uuid(),
    account_id   uuid not null unique references accounts (id),
    display_name text not null,
    status       text not null default 'created'
                 check (status in ('created', 'incomplete', 'ready', 'updated')),
    timezone     text,
    created_at   timestamptz not null default now(),
    updated_at   timestamptz not null default now(),
    deleted_at   timestamptz
);

create table preference_sets (
    id              uuid primary key default gen_random_uuid(),
    user_profile_id uuid not null references user_profiles (id),
    status          text not null default 'declared'
                    check (status in ('declared', 'active', 'revised', 'retired')),
    preferences     jsonb not null default '{}',
    valid_from      date not null default current_date,
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now()
);

create index preference_sets_profile_idx on preference_sets (user_profile_id)
    where status in ('active', 'declared');

-- Pantry Management
create table pantries (
    id              uuid primary key default gen_random_uuid(),
    user_profile_id uuid not null unique references user_profiles (id),
    status          text not null default 'active'
                    check (status in ('active', 'archived')),
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now(),
    deleted_at      timestamptz
);

create table pantry_items (
    id              uuid primary key default gen_random_uuid(),
    pantry_id       uuid not null references pantries (id),
    ingredient_name text not null,
    category        text not null,
    quantity        numeric not null check (quantity >= 0),
    unit            text not null default 'unit',
    expiry_date     date,
    status          text not null default 'available'
                    check (status in ('available', 'running_low', 'expiring_soon', 'consumed', 'removed')),
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now(),
    deleted_at      timestamptz
);

create index pantry_items_pantry_idx on pantry_items (pantry_id);
create index pantry_items_expiry_idx on pantry_items (expiry_date)
    where status not in ('consumed', 'removed');
create index pantry_items_category_idx on pantry_items (category);

-- Culinary Knowledge and Recipe Experience
create table recipes (
    id                uuid primary key default gen_random_uuid(),
    title             text not null,
    summary           text not null,
    servings          int not null check (servings > 0),
    prep_time_minutes int check (prep_time_minutes >= 0),
    cook_time_minutes int check (cook_time_minutes >= 0),
    ingredients       jsonb not null,
    instructions      jsonb not null,
    is_public         boolean not null default true,
    status            text not null default 'available'
                      check (status in ('available', 'revised', 'retired')),
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    deleted_at        timestamptz
);

create index recipes_title_idx on recipes (lower(title));
create index recipes_public_status_idx on recipes (is_public, status);

create table recipe_favorites (
    id              uuid primary key default gen_random_uuid(),
    recipe_id       uuid not null references recipes (id),
    user_profile_id uuid not null references user_profiles (id),
    status          text not null default 'active'
                    check (status in ('active', 'removed')),
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now(),
    deleted_at      timestamptz
);

create unique index recipe_favorites_profile_recipe_idx
    on recipe_favorites (user_profile_id, recipe_id) where status = 'active';

-- Meal Planning
create table meal_plans (
    id              uuid primary key default gen_random_uuid(),
    user_profile_id uuid not null references user_profiles (id),
    period_start    date not null,
    period_end      date not null check (period_end >= period_start),
    status          text not null default 'draft'
                    check (status in ('draft', 'planned', 'in_progress', 'completed', 'cancelled', 'revised')),
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now(),
    deleted_at      timestamptz
);

create index meal_plans_profile_period_idx on meal_plans (user_profile_id, period_start);

-- AI-Assisted Kitchen Decision Support (dependency of planned_meals below)
create table kitchen_recommendations (
    id                uuid primary key default gen_random_uuid(),
    user_profile_id   uuid not null references user_profiles (id),
    context_reference jsonb not null,
    rationale         text not null,
    confidence_statement text,
    status            text not null default 'requested'
                      check (status in ('requested', 'created', 'presented', 'accepted', 'rejected', 'superseded', 'unable_to_complete')),
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    deleted_at        timestamptz
);

create index kitchen_recommendations_profile_idx on kitchen_recommendations (user_profile_id);
create index kitchen_recommendations_status_idx on kitchen_recommendations (status);

create table recommendation_options (
    id                 uuid primary key default gen_random_uuid(),
    recommendation_id  uuid not null references kitchen_recommendations (id),
    recipe_id          uuid references recipes (id),
    position           int not null check (position > 0),
    rationale          text not null,
    status             text not null default 'proposed'
                       check (status in ('proposed', 'selected', 'rejected', 'superseded')),
    created_at         timestamptz not null default now(),
    updated_at         timestamptz not null default now(),
    deleted_at         timestamptz
);

create index recommendation_options_recommendation_idx on recommendation_options (recommendation_id);

create table planned_meals (
    id                        uuid primary key default gen_random_uuid(),
    meal_plan_id              uuid not null references meal_plans (id),
    meal_date                 date not null,
    meal_occasion             text not null default 'dinner',
    recipe_id                 uuid references recipes (id),
    recommendation_option_id  uuid references recommendation_options (id),
    status                    text not null default 'proposed'
                              check (status in ('proposed', 'planned', 'revised', 'removed', 'completed')),
    created_at                timestamptz not null default now(),
    updated_at                timestamptz not null default now(),
    deleted_at                timestamptz
);

create index planned_meals_plan_idx on planned_meals (meal_plan_id);
create index planned_meals_date_idx on planned_meals (meal_date, meal_occasion);

-- AI-Assisted Kitchen Decision Support: conversation retained while active
-- (M4-DEC-013: 30-day window, no raw prompts).
create table recommendation_conversations (
    id                uuid primary key default gen_random_uuid(),
    recommendation_id uuid not null unique references kitchen_recommendations (id),
    status            text not null default 'open'
                      check (status in ('open', 'completed', 'archived', 'deleted')),
    context_snapshot  jsonb,
    expires_at        timestamptz,
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    deleted_at        timestamptz
);

create index recommendation_conversations_expiry_idx on recommendation_conversations (expires_at);

-- Shopping Optimization
create table shopping_lists (
    id                uuid primary key default gen_random_uuid(),
    user_profile_id   uuid not null references user_profiles (id),
    meal_plan_id      uuid references meal_plans (id),
    recommendation_id uuid references kitchen_recommendations (id),
    title             text not null,
    status            text not null default 'draft'
                      check (status in ('draft', 'generated', 'reviewed', 'active', 'completed', 'cancelled', 'archived', 'revised')),
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now(),
    deleted_at        timestamptz
);

create index shopping_lists_profile_idx on shopping_lists (user_profile_id);

create table shopping_items (
    id              uuid primary key default gen_random_uuid(),
    shopping_list_id uuid not null references shopping_lists (id),
    ingredient_name text not null,
    quantity        numeric not null default 1 check (quantity >= 0),
    unit            text not null default 'unit',
    source          text not null default 'manual'
                    check (source in ('manual', 'generated')),
    pantry_item_id  uuid references pantry_items (id),
    status          text not null default 'open'
                    check (status in ('open', 'completed', 'removed')),
    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now(),
    deleted_at      timestamptz
);

create index shopping_items_list_idx on shopping_items (shopping_list_id);
create index shopping_items_status_idx on shopping_items (shopping_list_id, status);

-- +goose Down
drop table if exists shopping_items;
drop table if exists shopping_lists;
drop table if exists recommendation_conversations;
drop table if exists planned_meals;
drop table if exists recommendation_options;
drop table if exists kitchen_recommendations;
drop table if exists meal_plans;
drop table if exists recipe_favorites;
drop table if exists recipes;
drop table if exists pantry_items;
drop table if exists pantries;
drop table if exists preference_sets;
drop table if exists user_profiles;
drop table if exists auth_sessions;
drop table if exists accounts;
