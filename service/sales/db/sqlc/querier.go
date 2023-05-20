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
	DeletePOSItemsPOS(ctx context.Context, pointOfSaleID string) error
	GetPOS(ctx context.Context, arg GetPOSParams) ([]SalesPointOfSale, error)
	GetPOSCOASetting(ctx context.Context, branchID string) ([]SalesPosChartOfAccountSetting, error)
	GetPOSItemsByPOSId(ctx context.Context, pointOfSaleID string) ([]SalesPointOfSaleItem, error)
	GetPOSUserSetting(ctx context.Context, arg GetPOSUserSettingParams) (SalesPosUserSetting, error)
	InsertPOSCOASetting(ctx context.Context, arg InsertPOSCOASettingParams) (SalesPosChartOfAccountSetting, error)
	InsertPOSItem(ctx context.Context, arg InsertPOSItemParams) (SalesPointOfSaleItem, error)
	UpsertPOS(ctx context.Context, arg UpsertPOSParams) (SalesPointOfSale, error)
	UpsertPOSUserSetting(ctx context.Context, arg UpsertPOSUserSettingParams) (SalesPosUserSetting, error)
}

var _ Querier = (*Queries)(nil)
