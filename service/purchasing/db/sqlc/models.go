// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"
)

type PurchasingPurchaseOrder struct {
	ID                 string       `db:"id"`
	SalesOrderID       string       `db:"sales_order_id"`
	CompanyID          string       `db:"company_id"`
	BranchID           string       `db:"branch_id"`
	FormNumber         string       `db:"form_number"`
	TransactionDate    time.Time    `db:"transaction_date"`
	ContactBookID      string       `db:"contact_book_id"`
	SecondaryCompanyID string       `db:"secondary_company_id"`
	KonekinID          string       `db:"konekin_id"`
	CurrencyCode       string       `db:"currency_code"`
	TotalItems         int64        `db:"total_items"`
	Total              int64        `db:"total"`
	IsDeleted          bool         `db:"is_deleted"`
	Status             string       `db:"status"`
	CreatedAt          sql.NullTime `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
}

type PurchasingPurchaseOrderItem struct {
	ID                     string       `db:"id"`
	SalesOrderItemID       string       `db:"sales_order_item_id"`
	PurchaseOrderID        string       `db:"purchase_order_id"`
	PrimaryItemVariantID   string       `db:"primary_item_variant_id"`
	SecondaryItemVariantID string       `db:"secondary_item_variant_id"`
	PrimaryItemUnitID      string       `db:"primary_item_unit_id"`
	SecondaryItemUnitID    string       `db:"secondary_item_unit_id"`
	PrimaryItemUnitValue   int64        `db:"primary_item_unit_value"`
	SecondaryItemUnitValue int64        `db:"secondary_item_unit_value"`
	Amount                 int64        `db:"amount"`
	Price                  int64        `db:"price"`
	IsDeleted              bool         `db:"is_deleted"`
	CreatedAt              sql.NullTime `db:"created_at"`
	UpdatedAt              sql.NullTime `db:"updated_at"`
}