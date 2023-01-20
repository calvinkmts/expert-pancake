// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	AddCustomer(ctx context.Context, contactBookIds []string) error
	AddSupplier(ctx context.Context, contactBookIds []string) error
	DeleteContactBookBranches(ctx context.Context, contactBookID string) error
	GetContactBookAdditionalInfo(ctx context.Context, contactBookID string) (GetContactBookAdditionalInfoRow, error)
	GetContactBookBranches(ctx context.Context, contactBookID string) ([]GetContactBookBranchesRow, error)
	GetContactBookById(ctx context.Context, id string) (GetContactBookByIdRow, error)
	GetContactBookMailingAddress(ctx context.Context, contactBookID string) (GetContactBookMailingAddressRow, error)
	GetContactBookShippingAddress(ctx context.Context, contactBookID string) (GetContactBookShippingAddressRow, error)
	GetContactBooks(ctx context.Context, arg GetContactBooksParams) ([]GetContactBooksRow, error)
	GetContactGroups(ctx context.Context, companyID string) ([]GetContactGroupsRow, error)
	GetContactInvitations(ctx context.Context, companyID string) ([]GetContactInvitationsRow, error)
	GetCountKonekinId(ctx context.Context, konekinID string) (int64, error)
	GetCustomers(ctx context.Context, primaryCompanyID string) ([]GetCustomersRow, error)
	GetMyContactBook(ctx context.Context, primaryCompanyID string) (GetMyContactBookRow, error)
	GetReceiveInvitations(ctx context.Context, secondaryCompanyID string) ([]GetReceiveInvitationsRow, error)
	GetRequestInvitations(ctx context.Context, primaryCompanyID string) ([]GetRequestInvitationsRow, error)
	GetSuppliers(ctx context.Context, primaryCompanyID string) ([]GetSuppliersRow, error)
	InsertContactBook(ctx context.Context, arg InsertContactBookParams) (BusinessRelationContactBook, error)
	InsertContactBookAdditionalInfo(ctx context.Context, arg InsertContactBookAdditionalInfoParams) error
	InsertContactBookBranch(ctx context.Context, arg InsertContactBookBranchParams) error
	InsertContactBookMailingAddress(ctx context.Context, arg InsertContactBookMailingAddressParams) error
	InsertContactBookShippingAddress(ctx context.Context, arg InsertContactBookShippingAddressParams) error
	InsertContactGroup(ctx context.Context, arg InsertContactGroupParams) (BusinessRelationContactGroup, error)
	InsertContactInvitation(ctx context.Context, arg InsertContactInvitationParams) (BusinessRelationContactInvitation, error)
	UpdateContactBook(ctx context.Context, arg UpdateContactBookParams) (BusinessRelationContactBook, error)
	UpdateContactBookAdditionalInfo(ctx context.Context, arg UpdateContactBookAdditionalInfoParams) error
	UpdateContactBookGroupId(ctx context.Context, arg UpdateContactBookGroupIdParams) error
	UpdateContactBookGroupIdByGroupId(ctx context.Context, arg UpdateContactBookGroupIdByGroupIdParams) error
	UpdateContactBookMailingAddress(ctx context.Context, arg UpdateContactBookMailingAddressParams) error
	UpdateContactBookShippingAddress(ctx context.Context, arg UpdateContactBookShippingAddressParams) error
	UpdateContactBookTaxInfo(ctx context.Context, arg UpdateContactBookTaxInfoParams) error
	UpdateContactGroup(ctx context.Context, arg UpdateContactGroupParams) (BusinessRelationContactGroup, error)
	UpsertCustomerInfo(ctx context.Context, arg UpsertCustomerInfoParams) error
	UpsertSupplierInfo(ctx context.Context, arg UpsertSupplierInfoParams) error
}

var _ Querier = (*Queries)(nil)
