// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetItemBrands(ctx context.Context, arg GetItemBrandsParams) ([]GetItemBrandsRow, error)
	InsertItemBrand(ctx context.Context, arg InsertItemBrandParams) (InventoryItemBrand, error)
	InsertItemUnit(ctx context.Context, arg InsertItemUnitParams) (InventoryItemUnit, error)
	UpdateItemBrand(ctx context.Context, arg UpdateItemBrandParams) (InventoryItemBrand, error)
	UpdateItemUnit(ctx context.Context, arg UpdateItemUnitParams) (InventoryItemUnit, error)
}

var _ Querier = (*Queries)(nil)
