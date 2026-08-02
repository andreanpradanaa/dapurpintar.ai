-- Owning bounded context: Culinary Knowledge and Recipe Experience.
-- Business purpose: public recipe discovery, detail, and personal favorites.
-- Public queries filter is_public = true, status = 'available', deleted_at is null.
-- Favorites belong to a user profile; the partial unique index on
-- (user_profile_id, recipe_id) WHERE status = 'active' prevents duplicates.

-- name: ListPublicRecipes :many
select *
from recipes
where is_public = true
  and status = 'available'
  and deleted_at is null
  and (sqlc.narg('q')::text is null or lower(title) ilike '%' || sqlc.narg('q')::text || '%' or lower(summary) ilike '%' || sqlc.narg('q')::text || '%')
  and (sqlc.narg('max_prep')::int is null or prep_time_minutes is null or prep_time_minutes <= sqlc.narg('max_prep')::int)
  and ($1 = '' or id > $1::uuid)
order by
  case when sqlc.narg('sort_order')::text = 'title' then lower(title) end asc,
  case when sqlc.narg('sort_order')::text = 'prep_time' then prep_time_minutes end asc nulls last,
  case when sqlc.narg('sort_order')::text = '' then created_at end desc
limit $2;

-- name: GetRecipeByID :one
select *
from recipes
where id = $1
  and deleted_at is null
  and (is_public = true or sqlc.narg('include_private')::boolean = true);

-- name: GetActiveFavorite :one
select *
from recipe_favorites
where user_profile_id = $1
  and recipe_id = $2
  and status = 'active'
  and deleted_at is null;

-- name: CreateFavorite :one
insert into recipe_favorites (user_profile_id, recipe_id, status)
values ($1, $2, 'active')
on conflict (user_profile_id, recipe_id) where status = 'active' do nothing
returning *;

-- name: RemoveFavorite :one
update recipe_favorites
set status = 'removed', deleted_at = now(), updated_at = now()
where user_profile_id = $1
  and recipe_id = $2
  and status = 'active'
  and deleted_at is null
returning *;

-- name: ListFavorites :many
select rf.*, r.id as r_id, r.title as r_title, r.summary as r_summary,
       r.servings as r_servings, r.prep_time_minutes as r_prep_time_minutes,
       r.cook_time_minutes as r_cook_time_minutes, r.created_at as r_created_at
from recipe_favorites rf
join recipes r on r.id = rf.recipe_id and r.deleted_at is null
where rf.user_profile_id = $1
  and rf.status = 'active'
  and rf.deleted_at is null
  and ($2 = '' or rf.id > $2::uuid)
order by rf.created_at desc
limit $3;
