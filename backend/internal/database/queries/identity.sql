-- Owning bounded context: Identity and Access.
-- Business purpose: account lifecycle, authentication, and durable refresh-session
-- authority (M4-DEC-003). Queries filter soft-deleted rows unless named otherwise.

-- name: GetAccountByID :one
select *
from accounts
where id = $1
  and deleted_at is null;

-- name: GetAccountByEmail :one
select *
from accounts
where email = $1
  and deleted_at is null;

-- name: CreateAccount :one
insert into accounts (email, password_hash, status, timezone)
values ($1, $2, $3, $4)
returning *;

-- name: UpdateAccountStatus :one
update accounts
set status = $2, updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- name: MarkAccountEmailVerified :one
update accounts
set email_verified_at = now(), updated_at = now()
where id = $1
  and deleted_at is null
returning *;

-- Administrative only: includes soft-deleted rows.
-- name: ListAccountsIncludingDeleted :many
select *
from accounts
order by created_at desc;

-- name: CreateAuthSession :one
insert into auth_sessions (account_id, refresh_secret_hash, family_id, user_agent_hash, ip_hash, expires_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetAuthSessionBySecretHash :one
select *
from auth_sessions
where refresh_secret_hash = $1;

-- name: RevokeAuthSession :one
update auth_sessions
set revoked_at = now(), updated_at = now()
where id = $1
returning *;

-- name: RevokeSessionFamily :exec
update auth_sessions
set revoked_at = now(), updated_at = now()
where family_id = $1
  and revoked_at is null;

-- name: MarkAuthSessionReplacedBy :one
update auth_sessions
set replaced_by = $2, revoked_at = now(), updated_at = now()
where id = $1
returning *;

-- name: ListActiveSessionsForAccount :many
select *
from auth_sessions
where account_id = $1
  and revoked_at is null
  and expires_at > now();

-- name: DeleteExpiredAuthSessions :exec
delete from auth_sessions
where expires_at < now() - interval '7 days';
