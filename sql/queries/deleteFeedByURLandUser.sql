-- name: DeleteFeedRecord :exec
DELETE FROM feed_follows
WHERE user_id = $1 and feed_id = $2;
