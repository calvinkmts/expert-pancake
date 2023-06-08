// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	DeleteCompany(ctx context.Context, id string) error
	DeleteCompanyBranch(ctx context.Context, id string) error
	DeleteCompanyBranchesByCompanyId(ctx context.Context, companyID string) error
	GetCompanyBranchesByCompany(ctx context.Context, companyID string) ([]GetCompanyBranchesByCompanyRow, error)
	GetCompanyById(ctx context.Context, id string) (GetCompanyByIdRow, error)
	GetCompanyByName(ctx context.Context, name string) (string, error)
	GetUserCompanies(ctx context.Context, userID string) ([]GetUserCompaniesRow, error)
	GetUserCompaniesFilteredByName(ctx context.Context, arg GetUserCompaniesFilteredByNameParams) ([]GetUserCompaniesFilteredByNameRow, error)
	GetUserCompanyBranches(ctx context.Context, arg GetUserCompanyBranchesParams) ([]GetUserCompanyBranchesRow, error)
	GetUserCompanyBranchesFilteredByName(ctx context.Context, arg GetUserCompanyBranchesFilteredByNameParams) ([]GetUserCompanyBranchesFilteredByNameRow, error)
	InsertCompany(ctx context.Context, arg InsertCompanyParams) (BusinessCompany, error)
	InsertCompanyBranch(ctx context.Context, arg InsertCompanyBranchParams) (BusinessCompanyBranch, error)
	UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (BusinessCompany, error)
	UpdateCompanyBranch(ctx context.Context, arg UpdateCompanyBranchParams) (BusinessCompanyBranch, error)
}

var _ Querier = (*Queries)(nil)
