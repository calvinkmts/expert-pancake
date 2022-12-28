// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	DeleteChartOfAccountBranches(ctx context.Context, chartOfAccountID string) error
	GetCashTransactions(ctx context.Context, arg GetCashTransactionsParams) ([]GetCashTransactionsRow, error)
	GetCashTransactionsGroupByDate(ctx context.Context, arg GetCashTransactionsGroupByDateParams) ([]GetCashTransactionsGroupByDateRow, error)
	GetChartOfAccountBranches(ctx context.Context, chartOfAccountID string) ([]GetChartOfAccountBranchesRow, error)
	GetChartOfAccountGroup(ctx context.Context, id string) (GetChartOfAccountGroupRow, error)
	GetChartOfAccountGroupByAccTypeAccGroup(ctx context.Context, arg GetChartOfAccountGroupByAccTypeAccGroupParams) (GetChartOfAccountGroupByAccTypeAccGroupRow, error)
	GetChartOfAccountGroups(ctx context.Context, companyID string) ([]GetChartOfAccountGroupsRow, error)
	GetCompanyChartOfAccount(ctx context.Context, companyID string) (AccountingCompanyChartOfAccount, error)
	GetCompanyChartOfAccounts(ctx context.Context, arg GetCompanyChartOfAccountsParams) ([]GetCompanyChartOfAccountsRow, error)
	GetCompanySettingChartOfAccount(ctx context.Context, arg GetCompanySettingChartOfAccountParams) (GetCompanySettingChartOfAccountRow, error)
	GetCompanySettingFiscalYear(ctx context.Context, companyID string) (GetCompanySettingFiscalYearRow, error)
	InsertCashTransaction(ctx context.Context, arg InsertCashTransactionParams) (AccountingCashTransaction, error)
	InsertChartOfAccountBranches(ctx context.Context, arg InsertChartOfAccountBranchesParams) error
	InsertChartOfAccountGroup(ctx context.Context, arg InsertChartOfAccountGroupParams) (AccountingChartOfAccountGroup, error)
	InsertCompanyChartOfAccount(ctx context.Context, arg InsertCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error)
	InsertTransactionJournal(ctx context.Context, arg InsertTransactionJournalParams) (AccountingTransactionsJournal, error)
	UpdateChartOfAccountGroup(ctx context.Context, arg UpdateChartOfAccountGroupParams) (AccountingChartOfAccountGroup, error)
	UpdateCompanyChartOfAccount(ctx context.Context, arg UpdateCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error)
	UpsertCompanyFiscalYear(ctx context.Context, arg UpsertCompanyFiscalYearParams) (AccountingCompanyFiscalYear, error)
}

var _ Querier = (*Queries)(nil)
