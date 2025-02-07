-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;
