-- Owning bounded context: User Context and Preferences.
-- Business purpose: user profile and versioned preference sets.

-- name: GetUserProfileByAccountID :one
select *
from user_profiles
where account_id = $1
  and deleted_at is null;

-- name: GetUserProfileByID :one
select *
from user_profiles
where id = $1
  and deleted_at is null;

-- name: CreateUserProfile :one
insert into user_profiles (account_id, display_name, status, timezone)
values ($1, $2, $3, $4)
returning *;

-- name: UpdateUserProfile :one
update user_profiles
set display_name = coalesce($2, display_name),
    timezone = coalesce($3, timezone),
    status = 'updated',
    updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: GetActivePreferenceSetForProfile :one
select *
from preference_sets
where user_profile_id = $1
  and status in ('declared', 'active')
order by valid_from desc, created_at desc
limit 1;

-- name: CreatePreferenceSet :one
insert into preference_sets (user_profile_id, status, preferences, valid_from)
values ($1, 'active', $2, case when sqlc.arg('valid_from')::text = '' then current_date else sqlc.arg('valid_from')::date end)
returning *;

-- name: RetirePreferenceSet :exec
update preference_sets
set status = 'revised', updated_at = now()
where user_profile_id = $1
  and status in ('declared', 'active');
