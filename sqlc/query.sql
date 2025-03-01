-- name: GetCurrentVersion :one
SELECT version
FROM versions
WHERE id = ?;

-- name: UpdateVersion :exec
UPDATE versions
SET version = ?
WHERE id = ?;

-- name: InsertVersion :exec
INSERT INTO versions (
    id, version
) VALUES (
    ?, ?
);