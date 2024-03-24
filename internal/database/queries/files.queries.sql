-- name: GetFile :one
SELECT * FROM files
WHERE id = ? LIMIT 1;

-- name: CreateFile :execresult
INSERT INTO files (
  id, name, file_path
) VALUES (
  ?, ?, ?
);