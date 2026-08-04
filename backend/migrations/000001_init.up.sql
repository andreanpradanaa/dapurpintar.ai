CREATE TABLE IF NOT EXISTS recipes (
    id              TEXT PRIMARY KEY,
    slug            TEXT UNIQUE NOT NULL,
    title           TEXT NOT NULL,
    title_id        TEXT NOT NULL,
    description     TEXT NOT NULL,
    description_id  TEXT NOT NULL,
    image           TEXT,
    gradient        TEXT[] NOT NULL DEFAULT '{}',
    cuisine         TEXT NOT NULL,
    difficulty      TEXT NOT NULL,
    prep_time       INTEGER NOT NULL,
    cook_time       INTEGER NOT NULL,
    servings        INTEGER NOT NULL,
    ingredients     JSONB NOT NULL,
    steps           JSONB NOT NULL,
    nutrition       JSONB NOT NULL,
    tags            TEXT[] NOT NULL DEFAULT '{}',
    dietary         TEXT[] NOT NULL DEFAULT '{}',
    rating          REAL NOT NULL DEFAULT 0,
    reviews         INTEGER NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_recipes_slug ON recipes(slug);
CREATE INDEX IF NOT EXISTS idx_recipes_cuisine ON recipes(cuisine);
CREATE INDEX IF NOT EXISTS idx_recipes_dietary ON recipes USING GIN(dietary);
CREATE INDEX IF NOT EXISTS idx_recipes_tags ON recipes USING GIN(tags);
