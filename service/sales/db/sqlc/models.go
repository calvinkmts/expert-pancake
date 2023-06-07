// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
)

type SalesPointOfSale struct {
	ID                 string       `db:"id"`
	CompanyID          string       `db:"company_id"`
	BranchID           string       `db:"branch_id"`
	WarehouseID        string       `db:"warehouse_id"`
	FormNumber         string       `db:"form_number"`
	TransactionDate    time.Time    `db:"transaction_date"`
	ContactBookID      string       `db:"contact_book_id"`
	SecondaryCompanyID string       `db:"secondary_company_id"`
	KonekinID          string       `db:"konekin_id"`
	CurrencyCode       string       `db:"currency_code"`
	ChartOfAccountID   string       `db:"chart_of_account_id"`
	TotalItems         int64        `db:"total_items"`
	Total              int64        `db:"total"`
	IsDeleted          bool         `db:"is_deleted"`
	CreatedAt          sql.NullTime `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
}

type SalesPointOfSaleItem struct {
	ID              string         `db:"id"`
	PointOfSaleID   string         `db:"point_of_sale_id"`
	WarehouseRackID string         `db:"warehouse_rack_id"`
	ItemVariantID   string         `db:"item_variant_id"`
	ItemUnitID      string         `db:"item_unit_id"`
	ItemUnitValue   int64          `db:"item_unit_value"`
	Batch           sql.NullString `db:"batch"`
	ExpiredDate     sql.NullTime   `db:"expired_date"`
	ItemBarcodeID   string         `db:"item_barcode_id"`
	Amount          int64          `db:"amount"`
	Price           int64          `db:"price"`
	IsDeleted       bool           `db:"is_deleted"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	UpdatedAt       sql.NullTime   `db:"updated_at"`
}

type SalesPosChartOfAccountSetting struct {
	BranchID         string       `db:"branch_id"`
	ChartOfAccountID string       `db:"chart_of_account_id"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type SalesPosCustomerSetting struct {
	BranchID      string       `db:"branch_id"`
	ContactBookID string       `db:"contact_book_id"`
	CreatedAt     sql.NullTime `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}

type SalesPosPaymentMethod struct {
	ID               string       `db:"id"`
	CompanyID        string       `db:"company_id"`
	ChartOfAccountID string       `db:"chart_of_account_id"`
	Name             string       `db:"name"`
	IsDeleted        bool         `db:"is_deleted"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type SalesPosUserSetting struct {
	UserID          string       `db:"user_id"`
	BranchID        string       `db:"branch_id"`
	WarehouseID     string       `db:"warehouse_id"`
	WarehouseRackID string       `db:"warehouse_rack_id"`
	CreatedAt       sql.NullTime `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}
