-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: SaveUploadedVideo :one
INSERT INTO videos (user_id, type, file_path, file_name, size_in_bytes, duration, created_at) VALUES (?, ?, ?, ?, ?, ?, ?);