-- Owning bounded context: Pantry Management.
-- Business purpose: Pantry and Pantry Item CRUD, expiry attention view.
-- One pantry per user profile (M5-001: pantries.user_profile_id unique).
-- Items are owned by a pantry; every query filters deleted_at is null.
-- Cursor-based pagination uses the item id as the cursor key.

-- name: GetPantryByProfileID :one
select *
from pantries
where user_profile_id = $1
  and deleted_at is null;

-- name: GetPantryByID :one
select *
from pantries
where id = $1
  and deleted_at is null;

-- name: CreatePantry :one
insert into pantries (user_profile_id, status)
values ($1, 'active')
returning *;

-- name: GetPantryItemByID :one
select *
from pantry_items
where id = $1
  and deleted_at is null;

-- name: CreatePantryItem :one
insert into pantry_items (pantry_id, ingredient_name, category, quantity, unit, expiry_date, status)
values ($1, $2, $3, $4, $5, $6, 'available')
returning *;

-- name: UpdatePantryItem :one
update pantry_items
set quantity    = coalesce(sqlc.narg('quantity'),    quantity),
    unit        = coalesce(sqlc.narg('unit'),        unit),
    category    = coalesce(sqlc.narg('category'),    category),
    expiry_date = coalesce(sqlc.narg('expiry_date'), expiry_date),
    status      = coalesce(sqlc.narg('status'),      status),
    updated_at  = now()
where id = sqlc.arg('id')
  and deleted_at is null
returning *;

-- name: UpdatePantryItemStatus :one
update pantry_items
set status = $2, updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: RemovePantryItem :one
update pantry_items
set status = 'removed', deleted_at = now(), updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: RefreshPantryItemStatuses :exec
update pantry_items
set status = case
  when expiry_date is not null and expiry_date <= current_date + $2::int then 'expiring_soon'
  when quantity <= 1 then 'running_low'
  else status
end,
updated_at = now()
from pantries p
where p.id = pantry_items.pantry_id
  and p.user_profile_id = $1
  and pantry_items.status = 'available'
  and pantry_items.deleted_at is null;

-- name: ListPantryItems :many
select pi.*
from pantry_items pi
join pantries p on p.id = pi.pantry_id and p.deleted_at is null
where p.user_profile_id = $3
  and pi.deleted_at is null
  and (sqlc.narg('category')::text is null or pi.category = sqlc.narg('category'))
  and (sqlc.narg('status')::text is null or pi.status = sqlc.narg('status'))
  and ($1 = '' or pi.id > $1::uuid)
order by
  case when sqlc.narg('sort_order')::text = 'category'   then pi.category end asc,
  case when sqlc.narg('sort_order')::text = 'created_at' then pi.created_at end desc,
  case when sqlc.narg('sort_order')::text = ''           then pi.expiry_date end asc nulls last,
  case when sqlc.narg('sort_order')::text = 'expiry_date' then pi.expiry_date end asc nulls last
limit $2;

-- name: ListExpiringItems :many
select pi.*
from pantry_items pi
join pantries p on p.id = pi.pantry_id and p.deleted_at is null
where p.user_profile_id = $4
  and pi.deleted_at is null
  and pi.status in ('available', 'running_low', 'expiring_soon')
  and pi.expiry_date is not null
  and pi.expiry_date <= $3
  and ($1 = '' or pi.id > $1::uuid)
order by pi.expiry_date asc
limit $2;
