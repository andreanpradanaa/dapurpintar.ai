-- Owning bounded context: Meal Planning.
-- Business purpose: meal plan lifecycle and planned meal management.
-- Plans belong to a user profile; planned meals belong to a plan.
-- Slot conflict enforcement is at the application layer via the
-- planned_meals_date_idx index (meal_date, meal_occasion) within a plan.

-- name: CreateMealPlan :one
insert into meal_plans (user_profile_id, period_start, period_end, title, status)
values ($1, $2, $3, $4, 'draft')
returning *;

-- name: GetMealPlanByID :one
select *
from meal_plans
where id = $1
  and deleted_at is null;

-- name: ListMealPlans :many
select *
from meal_plans
where user_profile_id = $1
  and deleted_at is null
  and ($2 = '' or id > $2::uuid)
order by period_start desc
limit $3;

-- name: UpdateMealPlan :one
update meal_plans
set title        = coalesce(sqlc.narg('title'),        title),
    period_start = coalesce(sqlc.narg('period_start'), period_start),
    period_end   = coalesce(sqlc.narg('period_end'),   period_end),
    status       = 'revised',
    updated_at   = now()
where id = sqlc.arg('id')
  and deleted_at is null
returning *;

-- name: UpdateMealPlanStatus :one
update meal_plans
set status     = $2,
    updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: CreatePlannedMeal :one
insert into planned_meals (meal_plan_id, meal_date, meal_occasion, recipe_id, recommendation_option_id, status)
values ($1, $2, $3, $4, $5, 'planned')
returning *;

-- name: ListPlannedMeals :many
select pm.*
from planned_meals pm
join meal_plans mp on mp.id = pm.meal_plan_id and mp.deleted_at is null
where mp.user_profile_id = $1
  and pm.meal_plan_id = $2
  and pm.deleted_at is null
  and ($3 = '' or pm.id > $3::uuid)
order by pm.meal_date asc, pm.meal_occasion asc
limit $4;

-- name: GetPlannedMealByID :one
select *
from planned_meals
where id = $1
  and deleted_at is null;

-- name: UpdatePlannedMeal :one
update planned_meals
set meal_occasion = coalesce(sqlc.narg('occasion'), meal_occasion),
    recipe_id     = coalesce(sqlc.narg('recipe_id'), recipe_id),
    status        = coalesce(sqlc.narg('status'),    status),
    updated_at    = now()
where id = sqlc.arg('id')
  and deleted_at is null
returning *;

-- name: RemovePlannedMeal :one
update planned_meals
set status = 'removed', deleted_at = now(), updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: GetPlannedMealBySlot :one
select *
from planned_meals
where meal_plan_id = $1
  and meal_date = $2
  and meal_occasion = $3
  and status not in ('removed')
  and deleted_at is null;
