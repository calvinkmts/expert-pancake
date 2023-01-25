// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const getItemBrands = `-- name: GetItemBrands :many
SELECT id, company_id, name FROM inventory.item_brands
WHERE company_id = $1 AND name LIKE $2
`

type GetItemBrandsParams struct {
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetItemBrandsRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetItemBrands(ctx context.Context, arg GetItemBrandsParams) ([]GetItemBrandsRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemBrands, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemBrandsRow
	for rows.Next() {
		var i GetItemBrandsRow
		if err := rows.Scan(&i.ID, &i.CompanyID, &i.Name); err != nil {
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

const insertItemBrand = `-- name: InsertItemBrand :one
INSERT INTO inventory.item_brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING id, company_id, name, created_at, updated_at
`

type InsertItemBrandParams struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) InsertItemBrand(ctx context.Context, arg InsertItemBrandParams) (InventoryItemBrand, error) {
	row := q.db.QueryRowContext(ctx, insertItemBrand, arg.ID, arg.CompanyID, arg.Name)
	var i InventoryItemBrand
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateItemBrand = `-- name: UpdateItemBrand :one
UPDATE inventory.item_brands
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, name, created_at, updated_at
`

type UpdateItemBrandParams struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (q *Queries) UpdateItemBrand(ctx context.Context, arg UpdateItemBrandParams) (InventoryItemBrand, error) {
	row := q.db.QueryRowContext(ctx, updateItemBrand, arg.ID, arg.Name)
	var i InventoryItemBrand
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}