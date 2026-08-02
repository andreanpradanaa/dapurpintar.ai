-- Owning bounded context: AI-Assisted Kitchen Decision Support.
-- Business purpose: recommendation lifecycle and option management.
-- Recommendations belong to a user profile.

-- name: CreateRecommendation :one
insert into kitchen_recommendations (user_profile_id, context_reference, rationale, purpose, status)
values ($1, $2, '', $3, 'requested')
returning *;

-- name: GetRecommendationByID :one
select *
from kitchen_recommendations
where id = $1
  and deleted_at is null;

-- name: ListRecommendations :many
select *
from kitchen_recommendations
where user_profile_id = $1
  and deleted_at is null
  and (sqlc.narg('status')::text is null or status = sqlc.narg('status'))
  and (sqlc.narg('purpose')::text is null or purpose = sqlc.narg('purpose'))
  and ($2 = '' or id > $2::uuid)
order by created_at desc
limit $3;

-- name: UpdateRecommendationStatus :one
update kitchen_recommendations
set status     = $2,
    updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: CreateRecommendationOption :one
insert into recommendation_options (recommendation_id, recipe_id, position, rationale, status)
values ($1, $2, $3, $4, 'proposed')
returning *;

-- name: ListRecommendationOptions :many
select *
from recommendation_options
where recommendation_id = $1
  and deleted_at is null
order by position asc;

-- name: GetRecommendationOptionByID :one
select *
from recommendation_options
where id = $1
  and deleted_at is null;

-- name: UpdateRecommendationOptionStatus :one
update recommendation_options
set status     = $2,
    updated_at = now()
where id = $1
  and deleted_at is null
returning *;
