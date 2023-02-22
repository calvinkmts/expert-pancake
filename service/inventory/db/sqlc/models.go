// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type InventoryBrand struct {
	ID        string       `db:"id"`
	CompanyID string       `db:"company_id"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	IsDeleted bool         `db:"is_deleted"`
}

type InventoryGroup struct {
	ID        string       `db:"id"`
	CompanyID string       `db:"company_id"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	IsDeleted bool         `db:"is_deleted"`
}

type InventoryInternalStockTransfer struct {
	ID                     string       `db:"id"`
	SourceWarehouseID      string       `db:"source_warehouse_id"`
	DestinationWarehouseID string       `db:"destination_warehouse_id"`
	FormNumber             string       `db:"form_number"`
	TransactionDate        time.Time    `db:"transaction_date"`
	IsDeleted              bool         `db:"is_deleted"`
	CreatedAt              sql.NullTime `db:"created_at"`
	UpdatedAt              sql.NullTime `db:"updated_at"`
}

type InventoryInternalStockTransferItem struct {
	ID                      string         `db:"id"`
	InternalStockTransferID string         `db:"internal_stock_transfer_id"`
	WarehouseRackID         string         `db:"warehouse_rack_id"`
	VariantID               string         `db:"variant_id"`
	ItemUnitID              string         `db:"item_unit_id"`
	ItemUnitValue           int64          `db:"item_unit_value"`
	Amount                  int64          `db:"amount"`
	Batch                   sql.NullString `db:"batch"`
	ExpiredDate             sql.NullTime   `db:"expired_date"`
	ItemBarcodeID           string         `db:"item_barcode_id"`
	IsDeleted               bool           `db:"is_deleted"`
	CreatedAt               sql.NullTime   `db:"created_at"`
	UpdatedAt               sql.NullTime   `db:"updated_at"`
}

type InventoryItem struct {
	ID          string       `db:"id"`
	CompanyID   string       `db:"company_id"`
	ImageUrl    string       `db:"image_url"`
	Code        string       `db:"code"`
	Name        string       `db:"name"`
	BrandID     string       `db:"brand_id"`
	GroupID     string       `db:"group_id"`
	Tag         string       `db:"tag"`
	Description string       `db:"description"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type InventoryItemBarcode struct {
	ID          string         `db:"id"`
	VariantID   string         `db:"variant_id"`
	Batch       sql.NullString `db:"batch"`
	ExpiredDate sql.NullTime   `db:"expired_date"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
}

type InventoryItemInfo struct {
	ItemID                   string       `db:"item_id"`
	IsPurchase               bool         `db:"is_purchase"`
	IsSale                   bool         `db:"is_sale"`
	IsRawMaterial            bool         `db:"is_raw_material"`
	IsAsset                  bool         `db:"is_asset"`
	PurchaseChartOfAccountID string       `db:"purchase_chart_of_account_id"`
	SaleChartOfAccountID     string       `db:"sale_chart_of_account_id"`
	PurchaseItemUnitID       string       `db:"purchase_item_unit_id"`
	CreatedAt                sql.NullTime `db:"created_at"`
	UpdatedAt                sql.NullTime `db:"updated_at"`
}

type InventoryItemReorder struct {
	ID           string       `db:"id"`
	WarehouseID  string       `db:"warehouse_id"`
	ItemUnitID   string       `db:"item_unit_id"`
	VariantID    string       `db:"variant_id"`
	MinimumStock int64        `db:"minimum_stock"`
	CreatedAt    sql.NullTime `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}

type InventoryItemUnit struct {
	ID        string       `db:"id"`
	ItemID    string       `db:"item_id"`
	UnitID    string       `db:"unit_id"`
	Value     int64        `db:"value"`
	IsDefault bool         `db:"is_default"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type InventoryItemVariant struct {
	ID        string       `db:"id"`
	ItemID    string       `db:"item_id"`
	ImageUrl  string       `db:"image_url"`
	Barcode   string       `db:"barcode"`
	Name      string       `db:"name"`
	Price     int64        `db:"price"`
	IsDefault bool         `db:"is_default"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type InventoryStockMovement struct {
	ID                   string       `db:"id"`
	TransactionID        string       `db:"transaction_id"`
	TransactionDate      time.Time    `db:"transaction_date"`
	TransactionReference string       `db:"transaction_reference"`
	DetailTransactionID  string       `db:"detail_transaction_id"`
	WarehouseID          string       `db:"warehouse_id"`
	WarehouseRackID      string       `db:"warehouse_rack_id"`
	VariantID            string       `db:"variant_id"`
	ItemBarcodeID        string       `db:"item_barcode_id"`
	Amount               int64        `db:"amount"`
	CreatedAt            sql.NullTime `db:"created_at"`
	UpdatedAt            sql.NullTime `db:"updated_at"`
}

type InventoryUnit struct {
	ID             string       `db:"id"`
	UnitCategoryID string       `db:"unit_category_id"`
	CompanyID      string       `db:"company_id"`
	Name           string       `db:"name"`
	CreatedAt      sql.NullTime `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
}

type InventoryUnitCategory struct {
	ID        string       `db:"id"`
	CompanyID string       `db:"company_id"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type InventoryUpdateStock struct {
	ID              string         `db:"id"`
	FormNumber      string         `db:"form_number"`
	TransactionDate time.Time      `db:"transaction_date"`
	WarehouseID     string         `db:"warehouse_id"`
	WarehouseRackID string         `db:"warehouse_rack_id"`
	VariantID       string         `db:"variant_id"`
	ItemUnitID      string         `db:"item_unit_id"`
	ItemUnitValue   int64          `db:"item_unit_value"`
	BeginningStock  int64          `db:"beginning_stock"`
	EndingStock     int64          `db:"ending_stock"`
	Batch           sql.NullString `db:"batch"`
	ExpiredDate     sql.NullTime   `db:"expired_date"`
	ItemBarcodeID   string         `db:"item_barcode_id"`
	IsDeleted       bool           `db:"is_deleted"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	UpdatedAt       sql.NullTime   `db:"updated_at"`
}
