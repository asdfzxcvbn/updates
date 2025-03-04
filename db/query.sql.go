// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"
)

const getCurrentVersion = `-- name: GetCurrentVersion :one
SELECT version
FROM versions
WHERE id = ?
`

func (q *Queries) GetCurrentVersion(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getCurrentVersion, id)
	var version string
	err := row.Scan(&version)
	return version, err
}

const insertVersion = `-- name: InsertVersion :exec
INSERT INTO versions (
    id, version
) VALUES (
    ?, ?
)
`

type InsertVersionParams struct {
	ID      string
	Version string
}

func (q *Queries) InsertVersion(ctx context.Context, arg InsertVersionParams) error {
	_, err := q.db.ExecContext(ctx, insertVersion, arg.ID, arg.Version)
	return err
}

const updateVersion = `-- name: UpdateVersion :exec
UPDATE versions
SET version = ?
WHERE id = ?
`

type UpdateVersionParams struct {
	Version string
	ID      string
}

func (q *Queries) UpdateVersion(ctx context.Context, arg UpdateVersionParams) error {
	_, err := q.db.ExecContext(ctx, updateVersion, arg.Version, arg.ID)
	return err
}
