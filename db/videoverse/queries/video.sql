-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetVideoByID :one
SELECT *
FROM videos
WHERE id = ?;

-- name: GetVideosByUserID :many
SELECT *
FROM videos
WHERE user_id = ?;

-- name: GetLinksSharedByUserID :many
SELECT *
FROM shared_links
WHERE user_id = ?;

-- name: SaveUser :one
INSERT INTO users (email, username, password_hash, created_at)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: SaveVideo :one
INSERT INTO videos (user_id, title, description, type, file_path, file_name, size_in_bytes, duration, metadata,
                    created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: SaveSharedLink :one
INSERT INTO shared_links (user_id, video_id, link, expires_at, created_at)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteVideoByID :exec
DELETE
FROM videos
WHERE id = ?;
