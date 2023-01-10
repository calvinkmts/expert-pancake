// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	DeleteContactBookBranches(ctx context.Context, contactBookID string) error
	GetContactBookBranches(ctx context.Context, contactBookID string) ([]GetContactBookBranchesRow, error)
	GetContactBookById(ctx context.Context, id string) (GetContactBookByIdRow, error)
	GetContactBooks(ctx context.Context, primaryCompanyID string) ([]GetContactBooksRow, error)
	GetContactGroups(ctx context.Context, companyID string) ([]GetContactGroupsRow, error)
	InsertContactBook(ctx context.Context, arg InsertContactBookParams) (BusinessRelationContactBook, error)
	InsertContactBookAdditionalInfo(ctx context.Context, arg InsertContactBookAdditionalInfoParams) error
	InsertContactBookBranch(ctx context.Context, arg InsertContactBookBranchParams) error
	InsertContactBookMailingAddress(ctx context.Context, arg InsertContactBookMailingAddressParams) error
	InsertContactBookShippingAddress(ctx context.Context, arg InsertContactBookShippingAddressParams) error
	InsertContactGroup(ctx context.Context, arg InsertContactGroupParams) (BusinessRelationContactGroup, error)
	UpdateContactBook(ctx context.Context, arg UpdateContactBookParams) (BusinessRelationContactBook, error)
	UpdateContactBookAdditionalInfo(ctx context.Context, arg UpdateContactBookAdditionalInfoParams) error
	UpdateContactBookMailingAddress(ctx context.Context, arg UpdateContactBookMailingAddressParams) error
	UpdateContactBookShippingAddress(ctx context.Context, arg UpdateContactBookShippingAddressParams) error
	UpdateContactBookTaxInfo(ctx context.Context, arg UpdateContactBookTaxInfoParams) error
	UpdateContactGroup(ctx context.Context, arg UpdateContactGroupParams) (BusinessRelationContactGroup, error)
	UpsertCustomerInfo(ctx context.Context, arg UpsertCustomerInfoParams) error
}

var _ Querier = (*Queries)(nil)
