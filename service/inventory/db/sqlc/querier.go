// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	InsertItemBrand(ctx context.Context, arg InsertItemBrandParams) (InventoryItemBrand, error)
	UpdateItemBrand(ctx context.Context, arg UpdateItemBrandParams) (InventoryItemBrand, error)
}

var _ Querier = (*Queries)(nil)
