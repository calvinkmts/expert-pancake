// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type AccountingCashTransaction struct {
	ID                     string       `db:"id"`
	CompanyID              string       `db:"company_id"`
	BranchID               string       `db:"branch_id"`
	TransactionDate        time.Time    `db:"transaction_date"`
	Type                   string       `db:"type"`
	MainChartOfAccountID   string       `db:"main_chart_of_account_id"`
	ContraChartOfAccountID string       `db:"contra_chart_of_account_id"`
	Amount                 int64        `db:"amount"`
	Description            string       `db:"description"`
	CreatedAt              sql.NullTime `db:"created_at"`
	UpdatedAt              sql.NullTime `db:"updated_at"`
}

type AccountingChartOfAccountBranch struct {
	ChartOfAccountID string       `db:"chart_of_account_id"`
	BranchID         string       `db:"branch_id"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type AccountingChartOfAccountGroup struct {
	ID               string       `db:"id"`
	CompanyID        string       `db:"company_id"`
	ReportType       string       `db:"report_type"`
	AccountType      string       `db:"account_type"`
	AccountGroupName string       `db:"account_group_name"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type AccountingCompanyChartOfAccount struct {
	ID                    string       `db:"id"`
	CompanyID             string       `db:"company_id"`
	CurrencyCode          string       `db:"currency_code"`
	ChartOfAccountGroupID string       `db:"chart_of_account_group_id"`
	AccountCode           string       `db:"account_code"`
	AccountName           string       `db:"account_name"`
	BankName              string       `db:"bank_name"`
	BankAccountNumber     string       `db:"bank_account_number"`
	BankCode              string       `db:"bank_code"`
	IsDeleted             bool         `db:"is_deleted"`
	IsAllBranches         bool         `db:"is_all_branches"`
	CreatedAt             sql.NullTime `db:"created_at"`
	UpdatedAt             sql.NullTime `db:"updated_at"`
}

type AccountingJournalBook struct {
	ID          string       `db:"id"`
	CompanyID   string       `db:"company_id"`
	Name        string       `db:"name"`
	StartPeriod time.Time    `db:"start_period"`
	EndPeriod   time.Time    `db:"end_period"`
	IsClosed    bool         `db:"is_closed"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type AccountingJournalBookAccount struct {
	JournalBookID    string       `db:"journal_book_id"`
	ChartOfAccountID string       `db:"chart_of_account_id"`
	Amount           int64        `db:"amount"`
	CreatedAt        sql.NullTime `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
}

type AccountingMemorialJournal struct {
	ID              string       `db:"id"`
	CompanyID       string       `db:"company_id"`
	TransactionDate time.Time    `db:"transaction_date"`
	Description     string       `db:"description"`
	CreatedAt       sql.NullTime `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}

type AccountingMemorialJournalAccount struct {
	MemorialJournalID string       `db:"memorial_journal_id"`
	ChartOfAccountID  string       `db:"chart_of_account_id"`
	DebitAmount       int64        `db:"debit_amount"`
	CreditAmount      int64        `db:"credit_amount"`
	Description       string       `db:"description"`
	CreatedAt         sql.NullTime `db:"created_at"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
}

type AccountingTransactionsJournal struct {
	CompanyID            string       `db:"company_id"`
	BranchID             string       `db:"branch_id"`
	TransactionID        string       `db:"transaction_id"`
	TransactionDate      time.Time    `db:"transaction_date"`
	TransactionReference string       `db:"transaction_reference"`
	ChartOfAccountID     string       `db:"chart_of_account_id"`
	Amount               int64        `db:"amount"`
	Description          string       `db:"description"`
	CreatedAt            sql.NullTime `db:"created_at"`
}
