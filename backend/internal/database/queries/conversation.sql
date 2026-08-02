-- Owning bounded context: AI-Assisted Kitchen Decision Support.
-- Business purpose: recommendation-scoped conversation context (M4-DEC-013).
-- Messages are stored in context_snapshot JSONB.

-- name: CreateConversation :one
insert into recommendation_conversations (recommendation_id, context_snapshot, status)
values ($1, '[]', 'open')
returning *;

-- name: GetConversationByRecommendation :one
select *
from recommendation_conversations
where recommendation_id = $1
  and deleted_at is null;

-- name: UpdateConversationSnapshot :one
update recommendation_conversations
set context_snapshot = $2,
    updated_at       = now()
where recommendation_id = $1
  and deleted_at is null
returning *;

-- name: CloseConversation :one
update recommendation_conversations
set status     = 'completed',
    updated_at = now()
where recommendation_id = $1
  and status = 'open'
  and deleted_at is null
returning *;
