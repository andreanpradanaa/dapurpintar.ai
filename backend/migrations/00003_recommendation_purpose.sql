-- +goose Up
alter table kitchen_recommendations add column purpose text not null default '';

-- +goose Down
alter table kitchen_recommendations drop column purpose;
