-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    feeds.name as feed_name,
    users.name as user_name
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;
