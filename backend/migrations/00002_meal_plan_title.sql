-- +goose Up
alter table meal_plans add column title text not null default '';

-- +goose Down
alter table meal_plans drop column title;
