-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: DeleteFeeds :exec
DELETE FROM feeds;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2, updated_at = $3
WHERE id = $1;

-- name: GetNextFeedToFetch :one
select *
from feeds
order by last_fetched_at asc nulls first
limit 1;
