-- name: CreatePost :one
INSERT INTO posts(id,name,created_at,updated_at,url,feed_id,description,published_at,title)
 VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
JOIN feedsfollow ON posts.feed_id = feedsfollow.feed_id
WHERE feedsfollow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;

