// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetCompanyChartOfAccounts(ctx context.Context, arg GetCompanyChartOfAccountsParams) ([]GetCompanyChartOfAccountsRow, error)
}

var _ Querier = (*Queries)(nil)
