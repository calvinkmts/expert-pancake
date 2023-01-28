// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const upsertRack = `-- name: UpsertRack :one
INSERT INTO warehouse.racks(id, branch_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    branch_id = EXCLUDED.branch_id,
    updated_at = NOW()
RETURNING id, branch_id, name, created_at, updated_at
`

type UpsertRackParams struct {
	ID       string `db:"id"`
	BranchID string `db:"branch_id"`
	Name     string `db:"name"`
}

func (q *Queries) UpsertRack(ctx context.Context, arg UpsertRackParams) (WarehouseRack, error) {
	row := q.db.QueryRowContext(ctx, upsertRack, arg.ID, arg.BranchID, arg.Name)
	var i WarehouseRack
	err := row.Scan(
		&i.ID,
		&i.BranchID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}