-- name: CreateFeed :one
INSERT INTO feeds(name, url, user_id)
VALUES(
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeed :many
SELECT feeds.name, feeds.url, users.name FROM feeds
JOIN users ON feeds.user_id = users.id;
