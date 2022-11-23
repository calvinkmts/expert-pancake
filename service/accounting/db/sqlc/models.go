// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type AccountingCompanyChartOfAccount struct {
	ID                string       `db:"id"`
	CompanyID         string       `db:"company_id"`
	BranchID          string       `db:"branch_id"`
	AccountCode       string       `db:"account_code"`
	AccountName       string       `db:"account_name"`
	AccountGroup      string       `db:"account_group"`
	BankName          string       `db:"bank_name"`
	BankAccountNumber string       `db:"bank_account_number"`
	BankCode          string       `db:"bank_code"`
	OpeningBalance    float64      `db:"opening_balance"`
	IsDeleted         int32        `db:"is_deleted"`
	CreatedAt         sql.NullTime `db:"created_at"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
}

type AccountingCompanyFiscalYear struct {
	CompanyID  string       `db:"company_id"`
	StartMonth int32        `db:"start_month"`
	StartYear  int32        `db:"start_year"`
	EndMonth   int32        `db:"end_month"`
	EndYear    int32        `db:"end_year"`
	CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}
