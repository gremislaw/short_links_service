-- name: CreateShortURL :exec
INSERT INTO short_links (original_url, shortened_url)
VALUES ($1, $2);

-- name: GetOriginalURL :one
SELECT original_url
FROM short_links
WHERE shortened_url = $1;
