// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	UpsertRack(ctx context.Context, arg UpsertRackParams) (WarehouseRack, error)
}

var _ Querier = (*Queries)(nil)