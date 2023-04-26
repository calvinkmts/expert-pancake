// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetPurchaseOrderItems(ctx context.Context, purchaseOrderID string) ([]SalesPurchaseOrderItem, error)
	GetPurchaseOrders(ctx context.Context, arg GetPurchaseOrdersParams) ([]SalesPurchaseOrder, error)
	UpdatePurchaseOrderAddItem(ctx context.Context, purchaseOrderID string) error
	UpsertPurchaseOrder(ctx context.Context, arg UpsertPurchaseOrderParams) (SalesPurchaseOrder, error)
	UpsertPurchaseOrderItem(ctx context.Context, arg UpsertPurchaseOrderItemParams) (SalesPurchaseOrderItem, error)
}

var _ Querier = (*Queries)(nil)