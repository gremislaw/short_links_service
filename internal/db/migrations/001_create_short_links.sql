-- +goose Up


CREATE TABLE short_links (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    shortened_url TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индекс для поиска по shortened_url
CREATE INDEX idx_shortened_url ON short_links (shortened_url);

-- Индекс для поиска по original_url
CREATE INDEX idx_original_url ON short_links (original_url);

-- +goose Down

DROP INDEX idx_shortened_url;

DROP INDEX idx_original_url;

DROP TABLE short_links;