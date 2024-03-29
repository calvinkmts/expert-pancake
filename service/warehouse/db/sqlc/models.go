// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
)

type WarehouseWarehouse struct {
	ID        string       `db:"id"`
	BranchID  string       `db:"branch_id"`
	Code      string       `db:"code"`
	Name      string       `db:"name"`
	Address   string       `db:"address"`
	Type      string       `db:"type"`
	IsDeleted bool         `db:"is_deleted"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type WarehouseWarehouseRack struct {
	ID          string       `db:"id"`
	WarehouseID string       `db:"warehouse_id"`
	Name        string       `db:"name"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	IsDeleted   bool         `db:"is_deleted"`
}
