// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	GetUserCompanies(ctx context.Context, userID string) ([]GetUserCompaniesRow, error)
	GetUserCompaniesFilteredByName(ctx context.Context, arg GetUserCompaniesFilteredByNameParams) ([]GetUserCompaniesFilteredByNameRow, error)
	GetUserCompanyBranches(ctx context.Context, arg GetUserCompanyBranchesParams) ([]GetUserCompanyBranchesRow, error)
	GetUserCompanyBranchesFilteredByName(ctx context.Context, arg GetUserCompanyBranchesFilteredByNameParams) ([]GetUserCompanyBranchesFilteredByNameRow, error)
	UpsertCompany(ctx context.Context, arg UpsertCompanyParams) (BusinessCompany, error)
	UpsertCompanyBranch(ctx context.Context, arg UpsertCompanyBranchParams) (BusinessCompanyBranch, error)
}

var _ Querier = (*Queries)(nil)