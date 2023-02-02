// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const getWarehouses = `-- name: GetWarehouses :many
SELECT id, branch_id, code, name, address, type
FROM warehouse.warehouses
WHERE branch_id = $1 AND name LIKE $2
AND is_deleted = false
`

type GetWarehousesParams struct {
	BranchID string `db:"branch_id"`
	Name     string `db:"name"`
}

type GetWarehousesRow struct {
	ID       string `db:"id"`
	BranchID string `db:"branch_id"`
	Code     string `db:"code"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	Type     string `db:"type"`
}

func (q *Queries) GetWarehouses(ctx context.Context, arg GetWarehousesParams) ([]GetWarehousesRow, error) {
	rows, err := q.db.QueryContext(ctx, getWarehouses, arg.BranchID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWarehousesRow
	for rows.Next() {
		var i GetWarehousesRow
		if err := rows.Scan(
			&i.ID,
			&i.BranchID,
			&i.Code,
			&i.Name,
			&i.Address,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertWarehouse = `-- name: UpsertWarehouse :one
INSERT INTO warehouse.warehouses(id, branch_id, code, name, address, type)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id)
DO UPDATE SET
    branch_id = EXCLUDED.branch_id,
    name = EXCLUDED.name,
    address = EXCLUDED.address,
    type = EXCLUDED.type,
    updated_at = NOW()
RETURNING id, branch_id, code, name, address, type, is_deleted, created_at, updated_at
`

type UpsertWarehouseParams struct {
	ID       string `db:"id"`
	BranchID string `db:"branch_id"`
	Code     string `db:"code"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	Type     string `db:"type"`
}

func (q *Queries) UpsertWarehouse(ctx context.Context, arg UpsertWarehouseParams) (WarehouseWarehouse, error) {
	row := q.db.QueryRowContext(ctx, upsertWarehouse,
		arg.ID,
		arg.BranchID,
		arg.Code,
		arg.Name,
		arg.Address,
		arg.Type,
	)
	var i WarehouseWarehouse
	err := row.Scan(
		&i.ID,
		&i.BranchID,
		&i.Code,
		&i.Name,
		&i.Address,
		&i.Type,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
