-- name: GetFile :one
SELECT * FROM files
WHERE id = ? LIMIT 1;

-- name: CreateFile :execresult
INSERT INTO files (
  id, name, file_path, type
) VALUES (
  ?, ?, ?, ?
);

-- name: DeleteFile :exec
DELETE FROM files WHERE id = ?;