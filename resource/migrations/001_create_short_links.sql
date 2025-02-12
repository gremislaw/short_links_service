-- +goose Up

CREATE TABLE short_links (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    shortened_url TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE short_links;