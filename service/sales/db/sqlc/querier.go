// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	DeletePOS(ctx context.Context, id string) error
	DeletePOSCOASetting(ctx context.Context, branchID string) error
	DeletePOSCustomerSetting(ctx context.Context, branchID string) error
	DeletePOSItemsPOS(ctx context.Context, pointOfSaleID string) error
	DeletePOSPaymentMethod(ctx context.Context, id string) error
	DeleteSalesOrderItems(ctx context.Context, salesOrderID string) error
	GetCheckPOS(ctx context.Context, companyID string) (int64, error)
	GetPOS(ctx context.Context, arg GetPOSParams) ([]GetPOSRow, error)
	GetPOSCOASetting(ctx context.Context, branchID string) ([]SalesPosChartOfAccountSetting, error)
	GetPOSCustomerSetting(ctx context.Context, branchID string) ([]SalesPosCustomerSetting, error)
	GetPOSItemsByPOSId(ctx context.Context, pointOfSaleID string) ([]SalesPointOfSaleItem, error)
	GetPOSPaymentMethod(ctx context.Context, arg GetPOSPaymentMethodParams) ([]GetPOSPaymentMethodRow, error)
	GetPOSUserSetting(ctx context.Context, arg GetPOSUserSettingParams) (SalesPosUserSetting, error)
	InsertPOSCOASetting(ctx context.Context, arg InsertPOSCOASettingParams) (SalesPosChartOfAccountSetting, error)
	InsertPOSCustomerSetting(ctx context.Context, arg InsertPOSCustomerSettingParams) (SalesPosCustomerSetting, error)
	InsertPOSItem(ctx context.Context, arg InsertPOSItemParams) (SalesPointOfSaleItem, error)
	UpdateSalesOrderAddItem(ctx context.Context, salesOrderID string) error
	UpsertPOS(ctx context.Context, arg UpsertPOSParams) (SalesPointOfSale, error)
	UpsertPOSPaymentMethod(ctx context.Context, arg UpsertPOSPaymentMethodParams) error
	UpsertPOSUserSetting(ctx context.Context, arg UpsertPOSUserSettingParams) (SalesPosUserSetting, error)
	UpsertSalesOrder(ctx context.Context, arg UpsertSalesOrderParams) (SalesSalesOrder, error)
	UpsertSalesOrderItem(ctx context.Context, arg UpsertSalesOrderItemParams) (SalesSalesOrderItem, error)
}

var _ Querier = (*Queries)(nil)
