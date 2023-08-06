// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package db

import (
	"context"
)

const deleteCompany = `-- name: DeleteCompany :exec
UPDATE business.companies
SET is_deleted = true, 
updated_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, id)
	return err
}

const deleteCompanyBranch = `-- name: DeleteCompanyBranch :exec
UPDATE business.company_branches
SET is_deleted = true, 
updated_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCompanyBranch(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCompanyBranch, id)
	return err
}

const deleteCompanyBranchesByCompanyId = `-- name: DeleteCompanyBranchesByCompanyId :exec
UPDATE business.company_branches
SET is_deleted = true, 
updated_at = NOW()
WHERE company_id = $1
`

func (q *Queries) DeleteCompanyBranchesByCompanyId(ctx context.Context, companyID string) error {
	_, err := q.db.ExecContext(ctx, deleteCompanyBranchesByCompanyId, companyID)
	return err
}

const getCompanyBranches = `-- name: GetCompanyBranches :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE company_id = $1
ORDER BY is_central DESC
`

type GetCompanyBranchesRow struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	CompanyID   string `db:"company_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	IsCentral   bool   `db:"is_central"`
}

func (q *Queries) GetCompanyBranches(ctx context.Context, companyID string) ([]GetCompanyBranchesRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyBranches, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCompanyBranchesRow
	for rows.Next() {
		var i GetCompanyBranchesRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CompanyID,
			&i.Name,
			&i.Address,
			&i.PhoneNumber,
			&i.IsCentral,
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

const getCompanyBranchesByCompany = `-- name: GetCompanyBranchesByCompany :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE company_id = $1
`

type GetCompanyBranchesByCompanyRow struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	CompanyID   string `db:"company_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	IsCentral   bool   `db:"is_central"`
}

func (q *Queries) GetCompanyBranchesByCompany(ctx context.Context, companyID string) ([]GetCompanyBranchesByCompanyRow, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyBranchesByCompany, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCompanyBranchesByCompanyRow
	for rows.Next() {
		var i GetCompanyBranchesByCompanyRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CompanyID,
			&i.Name,
			&i.Address,
			&i.PhoneNumber,
			&i.IsCentral,
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

const getCompanyById = `-- name: GetCompanyById :one
SELECT id, user_id, name, initial_name, type, responsible_person
FROM business.companies
WHERE id = $1
`

type GetCompanyByIdRow struct {
	ID                string `db:"id"`
	UserID            string `db:"user_id"`
	Name              string `db:"name"`
	InitialName       string `db:"initial_name"`
	Type              string `db:"type"`
	ResponsiblePerson string `db:"responsible_person"`
}

func (q *Queries) GetCompanyById(ctx context.Context, id string) (GetCompanyByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getCompanyById, id)
	var i GetCompanyByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.InitialName,
		&i.Type,
		&i.ResponsiblePerson,
	)
	return i, err
}

const getCompanyByName = `-- name: GetCompanyByName :one
SELECT id FROM business.companies
WHERE name = $1
`

func (q *Queries) GetCompanyByName(ctx context.Context, name string) (string, error) {
	row := q.db.QueryRowContext(ctx, getCompanyByName, name)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getUserCompanies = `-- name: GetUserCompanies :many
SELECT id, user_id, name, initial_name, type, responsible_person, image_url FROM business.companies
WHERE user_id = $1 AND is_deleted = false
`

type GetUserCompaniesRow struct {
	ID                string `db:"id"`
	UserID            string `db:"user_id"`
	Name              string `db:"name"`
	InitialName       string `db:"initial_name"`
	Type              string `db:"type"`
	ResponsiblePerson string `db:"responsible_person"`
	ImageUrl          string `db:"image_url"`
}

func (q *Queries) GetUserCompanies(ctx context.Context, userID string) ([]GetUserCompaniesRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCompanies, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCompaniesRow
	for rows.Next() {
		var i GetUserCompaniesRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.InitialName,
			&i.Type,
			&i.ResponsiblePerson,
			&i.ImageUrl,
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

const getUserCompaniesFilteredByName = `-- name: GetUserCompaniesFilteredByName :many
SELECT id, user_id, name, initial_name, type, responsible_person, image_url FROM business.companies
WHERE user_id = $1 AND is_deleted = false AND name LIKE $2
`

type GetUserCompaniesFilteredByNameParams struct {
	UserID string `db:"user_id"`
	Name   string `db:"name"`
}

type GetUserCompaniesFilteredByNameRow struct {
	ID                string `db:"id"`
	UserID            string `db:"user_id"`
	Name              string `db:"name"`
	InitialName       string `db:"initial_name"`
	Type              string `db:"type"`
	ResponsiblePerson string `db:"responsible_person"`
	ImageUrl          string `db:"image_url"`
}

func (q *Queries) GetUserCompaniesFilteredByName(ctx context.Context, arg GetUserCompaniesFilteredByNameParams) ([]GetUserCompaniesFilteredByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCompaniesFilteredByName, arg.UserID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCompaniesFilteredByNameRow
	for rows.Next() {
		var i GetUserCompaniesFilteredByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.InitialName,
			&i.Type,
			&i.ResponsiblePerson,
			&i.ImageUrl,
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

const getUserCompanyBranches = `-- name: GetUserCompanyBranches :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE user_id = $1 AND company_id = $2 AND is_deleted = false
`

type GetUserCompanyBranchesParams struct {
	UserID    string `db:"user_id"`
	CompanyID string `db:"company_id"`
}

type GetUserCompanyBranchesRow struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	CompanyID   string `db:"company_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	IsCentral   bool   `db:"is_central"`
}

func (q *Queries) GetUserCompanyBranches(ctx context.Context, arg GetUserCompanyBranchesParams) ([]GetUserCompanyBranchesRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCompanyBranches, arg.UserID, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCompanyBranchesRow
	for rows.Next() {
		var i GetUserCompanyBranchesRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CompanyID,
			&i.Name,
			&i.Address,
			&i.PhoneNumber,
			&i.IsCentral,
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

const getUserCompanyBranchesFilteredByName = `-- name: GetUserCompanyBranchesFilteredByName :many
SELECT id, user_id, company_id, name, address, phone_number, is_central 
FROM business.company_branches
WHERE user_id = $1 AND company_id = $2 AND is_deleted = false AND name LIKE $3
`

type GetUserCompanyBranchesFilteredByNameParams struct {
	UserID    string `db:"user_id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

type GetUserCompanyBranchesFilteredByNameRow struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	CompanyID   string `db:"company_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	IsCentral   bool   `db:"is_central"`
}

func (q *Queries) GetUserCompanyBranchesFilteredByName(ctx context.Context, arg GetUserCompanyBranchesFilteredByNameParams) ([]GetUserCompanyBranchesFilteredByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserCompanyBranchesFilteredByName, arg.UserID, arg.CompanyID, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserCompanyBranchesFilteredByNameRow
	for rows.Next() {
		var i GetUserCompanyBranchesFilteredByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CompanyID,
			&i.Name,
			&i.Address,
			&i.PhoneNumber,
			&i.IsCentral,
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

const insertCompany = `-- name: InsertCompany :one
INSERT INTO business.companies(id, user_id, name, initial_name, type, responsible_person, image_url, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, user_id, name, initial_name, image_url, type, responsible_person, is_deleted, created_at, updated_at
`

type InsertCompanyParams struct {
	ID                string `db:"id"`
	UserID            string `db:"user_id"`
	Name              string `db:"name"`
	InitialName       string `db:"initial_name"`
	Type              string `db:"type"`
	ResponsiblePerson string `db:"responsible_person"`
	ImageUrl          string `db:"image_url"`
	IsDeleted         bool   `db:"is_deleted"`
}

func (q *Queries) InsertCompany(ctx context.Context, arg InsertCompanyParams) (BusinessCompany, error) {
	row := q.db.QueryRowContext(ctx, insertCompany,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.InitialName,
		arg.Type,
		arg.ResponsiblePerson,
		arg.ImageUrl,
		arg.IsDeleted,
	)
	var i BusinessCompany
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.InitialName,
		&i.ImageUrl,
		&i.Type,
		&i.ResponsiblePerson,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertCompanyBranch = `-- name: InsertCompanyBranch :one
INSERT INTO business.company_branches(id, user_id, company_id, name, address, phone_number, is_central, is_deleted)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, user_id, company_id, name, address, phone_number, is_central, is_deleted, created_at, updated_at
`

type InsertCompanyBranchParams struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	CompanyID   string `db:"company_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	IsCentral   bool   `db:"is_central"`
	IsDeleted   bool   `db:"is_deleted"`
}

func (q *Queries) InsertCompanyBranch(ctx context.Context, arg InsertCompanyBranchParams) (BusinessCompanyBranch, error) {
	row := q.db.QueryRowContext(ctx, insertCompanyBranch,
		arg.ID,
		arg.UserID,
		arg.CompanyID,
		arg.Name,
		arg.Address,
		arg.PhoneNumber,
		arg.IsCentral,
		arg.IsDeleted,
	)
	var i BusinessCompanyBranch
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CompanyID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
		&i.IsCentral,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertMemberRequest = `-- name: InsertMemberRequest :exec
INSERT INTO business.company_member_requests(
    id, user_id,company_id
)
VALUES ($1, $2, $3)
`

type InsertMemberRequestParams struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	CompanyID string `db:"company_id"`
}

func (q *Queries) InsertMemberRequest(ctx context.Context, arg InsertMemberRequestParams) error {
	_, err := q.db.ExecContext(ctx, insertMemberRequest, arg.ID, arg.UserID, arg.CompanyID)
	return err
}

const updateCompany = `-- name: UpdateCompany :one
UPDATE business.companies
SET name = $2, 
initial_name = $3, 
type = $4, 
responsible_person = $5, 
image_url = $6, 
updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, name, initial_name, image_url, type, responsible_person, is_deleted, created_at, updated_at
`

type UpdateCompanyParams struct {
	ID                string `db:"id"`
	Name              string `db:"name"`
	InitialName       string `db:"initial_name"`
	Type              string `db:"type"`
	ResponsiblePerson string `db:"responsible_person"`
	ImageUrl          string `db:"image_url"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (BusinessCompany, error) {
	row := q.db.QueryRowContext(ctx, updateCompany,
		arg.ID,
		arg.Name,
		arg.InitialName,
		arg.Type,
		arg.ResponsiblePerson,
		arg.ImageUrl,
	)
	var i BusinessCompany
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.InitialName,
		&i.ImageUrl,
		&i.Type,
		&i.ResponsiblePerson,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCompanyBranch = `-- name: UpdateCompanyBranch :one
UPDATE business.company_branches
SET name = $2, 
address = $3, 
phone_number = $4, 
updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, company_id, name, address, phone_number, is_central, is_deleted, created_at, updated_at
`

type UpdateCompanyBranchParams struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
}

func (q *Queries) UpdateCompanyBranch(ctx context.Context, arg UpdateCompanyBranchParams) (BusinessCompanyBranch, error) {
	row := q.db.QueryRowContext(ctx, updateCompanyBranch,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.PhoneNumber,
	)
	var i BusinessCompanyBranch
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CompanyID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
		&i.IsCentral,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
