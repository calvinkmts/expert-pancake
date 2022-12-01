// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getCashTransactions = `-- name: GetCashTransactions :many
SELECT a.id, a.company_id, a.branch_id, a.transaction_date, 
a.transaction_type, a.type, a.main_chart_of_account_id, a.contra_chart_of_account_id, 
a.amount, a.description, b.account_name AS main_chart_of_account_name, 
c.account_name AS contra_chart_of_account_name
FROM accounting.cash_transactions a
JOIN accounting.company_chart_of_accounts b ON a.main_chart_of_account_id = b.id
LEFT JOIN accounting.company_chart_of_accounts c ON a.contra_chart_of_account_id = c.id
WHERE a.company_id = $1
AND a.branch_id = $2 AND a.type LIKE $3
`

type GetCashTransactionsParams struct {
	CompanyID string `db:"company_id"`
	BranchID  string `db:"branch_id"`
	Type      string `db:"type"`
}

type GetCashTransactionsRow struct {
	ID                       string         `db:"id"`
	CompanyID                string         `db:"company_id"`
	BranchID                 string         `db:"branch_id"`
	TransactionDate          time.Time      `db:"transaction_date"`
	TransactionType          string         `db:"transaction_type"`
	Type                     string         `db:"type"`
	MainChartOfAccountID     string         `db:"main_chart_of_account_id"`
	ContraChartOfAccountID   string         `db:"contra_chart_of_account_id"`
	Amount                   int64          `db:"amount"`
	Description              string         `db:"description"`
	MainChartOfAccountName   string         `db:"main_chart_of_account_name"`
	ContraChartOfAccountName sql.NullString `db:"contra_chart_of_account_name"`
}

func (q *Queries) GetCashTransactions(ctx context.Context, arg GetCashTransactionsParams) ([]GetCashTransactionsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCashTransactions, arg.CompanyID, arg.BranchID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCashTransactionsRow
	for rows.Next() {
		var i GetCashTransactionsRow
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.BranchID,
			&i.TransactionDate,
			&i.TransactionType,
			&i.Type,
			&i.MainChartOfAccountID,
			&i.ContraChartOfAccountID,
			&i.Amount,
			&i.Description,
			&i.MainChartOfAccountName,
			&i.ContraChartOfAccountName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCashTransactionsGroupByDate = `-- name: GetCashTransactionsGroupByDate :many
SELECT transaction_date, 
SUM(CASE WHEN type = 'IN' THEN amount ELSE 0 END) AS cash_in, 
SUM(CASE WHEN type = 'OUT' THEN amount ELSE 0 END) AS cash_out
FROM accounting.cash_transactions 
WHERE company_id = $1
AND branch_id = $2
GROUP BY transaction_date
`

type GetCashTransactionsGroupByDateParams struct {
	CompanyID string `db:"company_id"`
	BranchID  string `db:"branch_id"`
}

type GetCashTransactionsGroupByDateRow struct {
	TransactionDate time.Time `db:"transaction_date"`
	CashIn          int64     `db:"cash_in"`
	CashOut         int64     `db:"cash_out"`
}

func (q *Queries) GetCashTransactionsGroupByDate(ctx context.Context, arg GetCashTransactionsGroupByDateParams) ([]GetCashTransactionsGroupByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCashTransactionsGroupByDate, arg.CompanyID, arg.BranchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCashTransactionsGroupByDateRow
	for rows.Next() {
		var i GetCashTransactionsGroupByDateRow
		if err := rows.Scan(&i.TransactionDate, &i.CashIn, &i.CashOut); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanyChartOfAccount = `-- name: GetCompanyChartOfAccount :one
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, bank_account_number, bank_code, opening_balance, is_deleted, created_at, updated_at
FROM accounting.company_chart_of_accounts
WHERE id = $1
`

func (q *Queries) GetCompanyChartOfAccount(ctx context.Context, id string) (AccountingCompanyChartOfAccount, error) {
	row := q.db.QueryRowContext(ctx, getCompanyChartOfAccount, id)
	var i AccountingCompanyChartOfAccount
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCompanyChartOfAccounts = `-- name: GetCompanyChartOfAccounts :many
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 AND account_name LIKE $2
AND account_group LIKE $3 
AND CASE WHEN $4 = '0' THEN is_deleted = FALSE 
WHEN $4 = '1' THEN is_deleted = TRUE ELSE 1=1 END
`

type GetCompanyChartOfAccountsParams struct {
	CompanyID    string      `db:"company_id"`
	AccountName  string      `db:"account_name"`
	AccountGroup string      `db:"account_group"`
	Column4      interface{} `db:"column_4"`
}

type GetCompanyChartOfAccountsRow struct {
	ID                string `db:"id"`
	CompanyID         string `db:"company_id"`
	BranchID          string `db:"branch_id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) GetCompanyChartOfAccounts(ctx context.Context, arg GetCompanyChartOfAccountsParams) ([]GetCompanyChartOfAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyChartOfAccounts,
		arg.CompanyID,
		arg.AccountName,
		arg.AccountGroup,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCompanyChartOfAccountsRow
	for rows.Next() {
		var i GetCompanyChartOfAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.BranchID,
			&i.AccountCode,
			&i.AccountName,
			&i.AccountGroup,
			&i.BankName,
			&i.BankAccountNumber,
			&i.BankCode,
			&i.OpeningBalance,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompanySettingBank = `-- name: GetCompanySettingBank :one
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 
AND account_group = 'BANK'
ORDER BY created_at LIMIT 1
`

type GetCompanySettingBankRow struct {
	ID                string `db:"id"`
	CompanyID         string `db:"company_id"`
	BranchID          string `db:"branch_id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) GetCompanySettingBank(ctx context.Context, companyID string) (GetCompanySettingBankRow, error) {
	row := q.db.QueryRowContext(ctx, getCompanySettingBank, companyID)
	var i GetCompanySettingBankRow
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
	)
	return i, err
}

const getCompanySettingCash = `-- name: GetCompanySettingCash :one
SELECT id, company_id, branch_id, account_code, account_name, account_group, bank_name, 
bank_account_number, bank_code, opening_balance, is_deleted
FROM accounting.company_chart_of_accounts
WHERE company_id = $1 
AND account_group = 'KAS'
ORDER BY created_at LIMIT 1
`

type GetCompanySettingCashRow struct {
	ID                string `db:"id"`
	CompanyID         string `db:"company_id"`
	BranchID          string `db:"branch_id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) GetCompanySettingCash(ctx context.Context, companyID string) (GetCompanySettingCashRow, error) {
	row := q.db.QueryRowContext(ctx, getCompanySettingCash, companyID)
	var i GetCompanySettingCashRow
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
	)
	return i, err
}

const getCompanySettingFiscalYear = `-- name: GetCompanySettingFiscalYear :one
SELECT company_id, start_period, end_period
FROM accounting.company_fiscal_years
WHERE company_id = $1
`

type GetCompanySettingFiscalYearRow struct {
	CompanyID   string    `db:"company_id"`
	StartPeriod time.Time `db:"start_period"`
	EndPeriod   time.Time `db:"end_period"`
}

func (q *Queries) GetCompanySettingFiscalYear(ctx context.Context, companyID string) (GetCompanySettingFiscalYearRow, error) {
	row := q.db.QueryRowContext(ctx, getCompanySettingFiscalYear, companyID)
	var i GetCompanySettingFiscalYearRow
	err := row.Scan(&i.CompanyID, &i.StartPeriod, &i.EndPeriod)
	return i, err
}

const insertCashTransaction = `-- name: InsertCashTransaction :one
INSERT INTO accounting.cash_transactions(id, company_id, branch_id, transaction_date, 
transaction_type, type, main_chart_of_account_id, contra_chart_of_account_id, 
amount, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, company_id, branch_id, transaction_date, transaction_type, type, main_chart_of_account_id, contra_chart_of_account_id, amount, description, created_at, updated_at
`

type InsertCashTransactionParams struct {
	ID                     string    `db:"id"`
	CompanyID              string    `db:"company_id"`
	BranchID               string    `db:"branch_id"`
	TransactionDate        time.Time `db:"transaction_date"`
	TransactionType        string    `db:"transaction_type"`
	Type                   string    `db:"type"`
	MainChartOfAccountID   string    `db:"main_chart_of_account_id"`
	ContraChartOfAccountID string    `db:"contra_chart_of_account_id"`
	Amount                 int64     `db:"amount"`
	Description            string    `db:"description"`
}

func (q *Queries) InsertCashTransaction(ctx context.Context, arg InsertCashTransactionParams) (AccountingCashTransaction, error) {
	row := q.db.QueryRowContext(ctx, insertCashTransaction,
		arg.ID,
		arg.CompanyID,
		arg.BranchID,
		arg.TransactionDate,
		arg.TransactionType,
		arg.Type,
		arg.MainChartOfAccountID,
		arg.ContraChartOfAccountID,
		arg.Amount,
		arg.Description,
	)
	var i AccountingCashTransaction
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.TransactionDate,
		&i.TransactionType,
		&i.Type,
		&i.MainChartOfAccountID,
		&i.ContraChartOfAccountID,
		&i.Amount,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertCompanyChartOfAccount = `-- name: InsertCompanyChartOfAccount :one
INSERT INTO accounting.company_chart_of_accounts(id, company_id, branch_id, 
account_code, account_name, account_group, 
bank_name, bank_account_number, bank_code, opening_balance, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, company_id, branch_id, account_code, account_name, account_group, bank_name, bank_account_number, bank_code, opening_balance, is_deleted, created_at, updated_at
`

type InsertCompanyChartOfAccountParams struct {
	ID                string `db:"id"`
	CompanyID         string `db:"company_id"`
	BranchID          string `db:"branch_id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) InsertCompanyChartOfAccount(ctx context.Context, arg InsertCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error) {
	row := q.db.QueryRowContext(ctx, insertCompanyChartOfAccount,
		arg.ID,
		arg.CompanyID,
		arg.BranchID,
		arg.AccountCode,
		arg.AccountName,
		arg.AccountGroup,
		arg.BankName,
		arg.BankAccountNumber,
		arg.BankCode,
		arg.OpeningBalance,
		arg.IsDeleted,
	)
	var i AccountingCompanyChartOfAccount
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertTransactionJournal = `-- name: InsertTransactionJournal :one
INSERT INTO accounting.transactions_journal(company_id, branch_id, transaction_id, 
transaction_date, transaction_reference , transaction_type, chart_of_account_id, 
amount, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING company_id, branch_id, transaction_id, transaction_date, transaction_reference, transaction_type, chart_of_account_id, amount, description, created_at
`

type InsertTransactionJournalParams struct {
	CompanyID            string    `db:"company_id"`
	BranchID             string    `db:"branch_id"`
	TransactionID        string    `db:"transaction_id"`
	TransactionDate      time.Time `db:"transaction_date"`
	TransactionReference string    `db:"transaction_reference"`
	TransactionType      string    `db:"transaction_type"`
	ChartOfAccountID     string    `db:"chart_of_account_id"`
	Amount               int64     `db:"amount"`
	Description          string    `db:"description"`
}

func (q *Queries) InsertTransactionJournal(ctx context.Context, arg InsertTransactionJournalParams) (AccountingTransactionsJournal, error) {
	row := q.db.QueryRowContext(ctx, insertTransactionJournal,
		arg.CompanyID,
		arg.BranchID,
		arg.TransactionID,
		arg.TransactionDate,
		arg.TransactionReference,
		arg.TransactionType,
		arg.ChartOfAccountID,
		arg.Amount,
		arg.Description,
	)
	var i AccountingTransactionsJournal
	err := row.Scan(
		&i.CompanyID,
		&i.BranchID,
		&i.TransactionID,
		&i.TransactionDate,
		&i.TransactionReference,
		&i.TransactionType,
		&i.ChartOfAccountID,
		&i.Amount,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const updateCompanyChartOfAccount = `-- name: UpdateCompanyChartOfAccount :one
UPDATE accounting.company_chart_of_accounts
SET 
    account_code = $2,
    account_name = $3,
    account_group = $4,
    bank_name = $5,
    bank_account_number = $6,
    bank_code = $7,
    opening_balance = $8,
    is_deleted = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, branch_id, account_code, account_name, account_group, bank_name, bank_account_number, bank_code, opening_balance, is_deleted, created_at, updated_at
`

type UpdateCompanyChartOfAccountParams struct {
	ID                string `db:"id"`
	AccountCode       string `db:"account_code"`
	AccountName       string `db:"account_name"`
	AccountGroup      string `db:"account_group"`
	BankName          string `db:"bank_name"`
	BankAccountNumber string `db:"bank_account_number"`
	BankCode          string `db:"bank_code"`
	OpeningBalance    int64  `db:"opening_balance"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) UpdateCompanyChartOfAccount(ctx context.Context, arg UpdateCompanyChartOfAccountParams) (AccountingCompanyChartOfAccount, error) {
	row := q.db.QueryRowContext(ctx, updateCompanyChartOfAccount,
		arg.ID,
		arg.AccountCode,
		arg.AccountName,
		arg.AccountGroup,
		arg.BankName,
		arg.BankAccountNumber,
		arg.BankCode,
		arg.OpeningBalance,
		arg.IsDeleted,
	)
	var i AccountingCompanyChartOfAccount
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.BranchID,
		&i.AccountCode,
		&i.AccountName,
		&i.AccountGroup,
		&i.BankName,
		&i.BankAccountNumber,
		&i.BankCode,
		&i.OpeningBalance,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertCompanyFiscalYear = `-- name: UpsertCompanyFiscalYear :one
INSERT INTO accounting.company_fiscal_years(company_id, start_period, end_period)
VALUES ($1, $2, $3)
ON CONFLICT (company_id)
DO UPDATE SET
    start_period = EXCLUDED.start_period,
    end_period = EXCLUDED.end_period,
    updated_at = NOW()
RETURNING company_id, start_period, end_period, created_at, updated_at
`

type UpsertCompanyFiscalYearParams struct {
	CompanyID   string    `db:"company_id"`
	StartPeriod time.Time `db:"start_period"`
	EndPeriod   time.Time `db:"end_period"`
}

func (q *Queries) UpsertCompanyFiscalYear(ctx context.Context, arg UpsertCompanyFiscalYearParams) (AccountingCompanyFiscalYear, error) {
	row := q.db.QueryRowContext(ctx, upsertCompanyFiscalYear, arg.CompanyID, arg.StartPeriod, arg.EndPeriod)
	var i AccountingCompanyFiscalYear
	err := row.Scan(
		&i.CompanyID,
		&i.StartPeriod,
		&i.EndPeriod,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
