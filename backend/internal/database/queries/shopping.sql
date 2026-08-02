-- Owning bounded context: Shopping Optimization.
-- Business purpose: shopping list lifecycle and item management.
-- Lists belong to a user profile; items belong to a list.
-- Source enum: manual | generated.

-- name: CreateShoppingList :one
insert into shopping_lists (user_profile_id, title, status)
values ($1, $2, $3)
returning *;

-- name: GetShoppingListByID :one
select *
from shopping_lists
where id = $1
  and deleted_at is null;

-- name: ListShoppingLists :many
select *
from shopping_lists
where user_profile_id = $1
  and deleted_at is null
  and ($2 = '' or id > $2::uuid)
order by created_at desc
limit $3;

-- name: UpdateShoppingList :one
update shopping_lists
set title      = coalesce(sqlc.narg('title'), title),
    status     = coalesce(sqlc.narg('status'), status),
    updated_at = now()
where id = sqlc.arg('id')
  and deleted_at is null
returning *;

-- name: UpdateShoppingListStatus :one
update shopping_lists
set status     = $2,
    updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: CreateShoppingItem :one
insert into shopping_items (shopping_list_id, ingredient_name, quantity, unit, source, status)
values ($1, $2, $3, $4, $5, 'open')
returning *;

-- name: ListShoppingItems :many
select si.*
from shopping_items si
join shopping_lists sl on sl.id = si.shopping_list_id and sl.deleted_at is null
where sl.user_profile_id = $1
  and si.shopping_list_id = $2
  and si.deleted_at is null
  and ($3 = '' or si.id > $3::uuid)
order by si.created_at asc
limit $4;

-- name: GetShoppingItemByID :one
select *
from shopping_items
where id = $1
  and deleted_at is null;

-- name: UpdateShoppingItem :one
update shopping_items
set ingredient_name = coalesce(sqlc.narg('ingredient_name'), ingredient_name),
    quantity        = coalesce(sqlc.narg('quantity'),        quantity),
    unit            = coalesce(sqlc.narg('unit'),            unit),
    updated_at      = now()
where id = sqlc.arg('id')
  and deleted_at is null
returning *;

-- name: UpdateShoppingItemStatus :one
update shopping_items
set status     = $2,
    updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: RemoveShoppingItem :one
update shopping_items
set status = 'removed', deleted_at = now(), updated_at = now()
where id = $1
  and deleted_at is null
returning *;
