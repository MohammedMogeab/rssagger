-- name: CreateFeedFollow :one
INSERT INTO feedsfollow(id,created_at,updated_at,feed_id,user_id)
 VALUES($1,$2,$3,$4,$5
)
RETURNING *;


-- name: GetFeedFollowsByUserID :many
SELECT * FROM feedsfollow WHERE user_id=$1;


-- name: DeleteFeedFollowByID :exec
DELETE FROM feedsfollow WHERE id=$1 AND user_id=$2;