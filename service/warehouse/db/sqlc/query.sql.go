// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const getWarehouseRacks = `-- name: GetWarehouseRacks :many
SELECT id, warehouse_id, name
FROM warehouse.warehouse_racks
WHERE warehouse_id = $1 AND name LIKE $2
`

type GetWarehouseRacksParams struct {
	WarehouseID string `db:"warehouse_id"`
	Name        string `db:"name"`
}

type GetWarehouseRacksRow struct {
	ID          string `db:"id"`
	WarehouseID string `db:"warehouse_id"`
	Name        string `db:"name"`
}

func (q *Queries) GetWarehouseRacks(ctx context.Context, arg GetWarehouseRacksParams) ([]GetWarehouseRacksRow, error) {
	rows, err := q.db.QueryContext(ctx, getWarehouseRacks, arg.WarehouseID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWarehouseRacksRow
	for rows.Next() {
		var i GetWarehouseRacksRow
		if err := rows.Scan(&i.ID, &i.WarehouseID, &i.Name); err != nil {
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

const upsertWarehouseRack = `-- name: UpsertWarehouseRack :one
INSERT INTO warehouse.warehouse_racks(id, warehouse_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    warehouse_id = EXCLUDED.warehouse_id,
    name = EXCLUDED.name,
    updated_at = NOW()
RETURNING id, warehouse_id, name, created_at, updated_at
`

type UpsertWarehouseRackParams struct {
	ID          string `db:"id"`
	WarehouseID string `db:"warehouse_id"`
	Name        string `db:"name"`
}

func (q *Queries) UpsertWarehouseRack(ctx context.Context, arg UpsertWarehouseRackParams) (WarehouseWarehouseRack, error) {
	row := q.db.QueryRowContext(ctx, upsertWarehouseRack, arg.ID, arg.WarehouseID, arg.Name)
	var i WarehouseWarehouseRack
	err := row.Scan(
		&i.ID,
		&i.WarehouseID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
