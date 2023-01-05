// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type BusinessCompany struct {
	ID                string       `db:"id"`
	UserID            string       `db:"user_id"`
	Name              string       `db:"name"`
	InitialName       string       `db:"initial_name"`
	Type              string       `db:"type"`
	ResponsiblePerson string       `db:"responsible_person"`
	IsDeleted         bool         `db:"is_deleted"`
	CreatedAt         sql.NullTime `db:"created_at"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
}

type BusinessCompanyBranch struct {
	ID          string       `db:"id"`
	UserID      string       `db:"user_id"`
	CompanyID   string       `db:"company_id"`
	Name        string       `db:"name"`
	Address     string       `db:"address"`
	PhoneNumber string       `db:"phone_number"`
	IsCentral   bool         `db:"is_central"`
	IsDeleted   bool         `db:"is_deleted"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type BusinessContactGroup struct {
	ID        string       `db:"id"`
	CompanyID string       `db:"company_id"`
	Name      string       `db:"name"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
