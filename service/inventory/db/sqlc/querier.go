// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetBrandById(ctx context.Context, id string) (GetBrandByIdRow, error)
	GetBrands(ctx context.Context, arg GetBrandsParams) ([]GetBrandsRow, error)
	GetGroupById(ctx context.Context, id string) (GetGroupByIdRow, error)
	GetGroups(ctx context.Context, arg GetGroupsParams) ([]GetGroupsRow, error)
	GetUnits(ctx context.Context, arg GetUnitsParams) ([]GetUnitsRow, error)
	InsertBrand(ctx context.Context, arg InsertBrandParams) (InventoryBrand, error)
	InsertGroup(ctx context.Context, arg InsertGroupParams) (InventoryGroup, error)
	InsertItem(ctx context.Context, arg InsertItemParams) (InventoryItem, error)
	InsertItemVariant(ctx context.Context, arg InsertItemVariantParams) (InventoryItemVariant, error)
	InsertUnit(ctx context.Context, arg InsertUnitParams) (InventoryUnit, error)
	UpdateBrand(ctx context.Context, arg UpdateBrandParams) (InventoryBrand, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (InventoryGroup, error)
	UpdateItem(ctx context.Context, arg UpdateItemParams) (InventoryItem, error)
	UpdateItemVariantDefault(ctx context.Context, arg UpdateItemVariantDefaultParams) (InventoryItemVariant, error)
	UpdateUnit(ctx context.Context, arg UpdateUnitParams) (InventoryUnit, error)
}

var _ Querier = (*Queries)(nil)
