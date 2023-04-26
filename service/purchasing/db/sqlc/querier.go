// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	DeletePurchaseOrderItems(ctx context.Context, purchaseOrderID string) error
	GetPurchaseOrderItems(ctx context.Context, purchaseOrderID string) ([]PurchasingPurchaseOrderItem, error)
	GetPurchaseOrders(ctx context.Context, arg GetPurchaseOrdersParams) ([]PurchasingPurchaseOrder, error)
	GetPurchaseSetting(ctx context.Context, companyID string) (PurchasingPurchaseSetting, error)
	UpdatePurchaseOrderAddItem(ctx context.Context, purchaseOrderID string) error
	UpsertPurchaseOrder(ctx context.Context, arg UpsertPurchaseOrderParams) (PurchasingPurchaseOrder, error)
	UpsertPurchaseOrderItem(ctx context.Context, arg UpsertPurchaseOrderItemParams) (PurchasingPurchaseOrderItem, error)
	UpsertPurchaseSetting(ctx context.Context, arg UpsertPurchaseSettingParams) (PurchasingPurchaseSetting, error)
}

var _ Querier = (*Queries)(nil)
