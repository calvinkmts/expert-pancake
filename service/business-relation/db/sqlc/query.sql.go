// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const deleteContactBookBranches = `-- name: DeleteContactBookBranches :exec
DELETE FROM business_relation.contact_book_branches
WHERE contact_book_id = $1
`

func (q *Queries) DeleteContactBookBranches(ctx context.Context, contactBookID string) error {
	_, err := q.db.ExecContext(ctx, deleteContactBookBranches, contactBookID)
	return err
}

const getContactBookBranches = `-- name: GetContactBookBranches :many
SELECT contact_book_id, company_branch_id FROM business_relation.contact_book_branches
WHERE contact_book_id = $1
`

type GetContactBookBranchesRow struct {
	ContactBookID   string `db:"contact_book_id"`
	CompanyBranchID string `db:"company_branch_id"`
}

func (q *Queries) GetContactBookBranches(ctx context.Context, contactBookID string) ([]GetContactBookBranchesRow, error) {
	rows, err := q.db.QueryContext(ctx, getContactBookBranches, contactBookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetContactBookBranchesRow
	for rows.Next() {
		var i GetContactBookBranchesRow
		if err := rows.Scan(&i.ContactBookID, &i.CompanyBranchID); err != nil {
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

const getContactBookById = `-- name: GetContactBookById :one
SELECT a.id, a.primary_company_id, a.secondary_company_id,
a.contact_group_id, b.name AS contact_group_name, a.name, a.email,
a.phone, a.mobile, a.web, a.is_all_branches, a.is_customer, a.is_supplier,
a.is_tax, a.tax_id, a.is_deleted
FROM business_relation.contact_books a
JOIN business_relation.contact_groups b ON a.contact_group_id = b.id
WHERE a.id = $1
`

type GetContactBookByIdRow struct {
	ID                 string `db:"id"`
	PrimaryCompanyID   string `db:"primary_company_id"`
	SecondaryCompanyID string `db:"secondary_company_id"`
	ContactGroupID     string `db:"contact_group_id"`
	ContactGroupName   string `db:"contact_group_name"`
	Name               string `db:"name"`
	Email              string `db:"email"`
	Phone              string `db:"phone"`
	Mobile             string `db:"mobile"`
	Web                string `db:"web"`
	IsAllBranches      bool   `db:"is_all_branches"`
	IsCustomer         bool   `db:"is_customer"`
	IsSupplier         bool   `db:"is_supplier"`
	IsTax              bool   `db:"is_tax"`
	TaxID              string `db:"tax_id"`
	IsDeleted          bool   `db:"is_deleted"`
}

func (q *Queries) GetContactBookById(ctx context.Context, id string) (GetContactBookByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getContactBookById, id)
	var i GetContactBookByIdRow
	err := row.Scan(
		&i.ID,
		&i.PrimaryCompanyID,
		&i.SecondaryCompanyID,
		&i.ContactGroupID,
		&i.ContactGroupName,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Mobile,
		&i.Web,
		&i.IsAllBranches,
		&i.IsCustomer,
		&i.IsSupplier,
		&i.IsTax,
		&i.TaxID,
		&i.IsDeleted,
	)
	return i, err
}

const getContactBooks = `-- name: GetContactBooks :many
SELECT a.id, a.primary_company_id, a.secondary_company_id, 
a.contact_group_id, a.name, a.email, a.phone, a.mobile, a.web,
a.is_all_branches, a.is_customer, a.is_supplier,
b.nickname, b.tag, b.note,
c.province AS mailing_province, c.regency AS mailing_regency,
c.district AS mailing_district, c.postal_code AS mailing_postal_code,
c.full_address AS mailing_full_address,
d.province AS shipping_province, d.regency AS shipping_regency,
d.district AS shipping_district, d.postal_code AS shipping_postal_code,
d.full_address AS shipping_full_address
FROM business_relation.contact_books a
JOIN business_relation.contact_book_additional_infos b ON a.id = b.contact_book_id
JOIN business_relation.contact_book_mailing_addresses c ON a.id = c.contact_book_id
JOIN business_relation.contact_book_shipping_addresses d ON a.id = d.contact_book_id
WHERE a.primary_company_id = $1
`

type GetContactBooksRow struct {
	ID                  string `db:"id"`
	PrimaryCompanyID    string `db:"primary_company_id"`
	SecondaryCompanyID  string `db:"secondary_company_id"`
	ContactGroupID      string `db:"contact_group_id"`
	Name                string `db:"name"`
	Email               string `db:"email"`
	Phone               string `db:"phone"`
	Mobile              string `db:"mobile"`
	Web                 string `db:"web"`
	IsAllBranches       bool   `db:"is_all_branches"`
	IsCustomer          bool   `db:"is_customer"`
	IsSupplier          bool   `db:"is_supplier"`
	Nickname            string `db:"nickname"`
	Tag                 string `db:"tag"`
	Note                string `db:"note"`
	MailingProvince     string `db:"mailing_province"`
	MailingRegency      string `db:"mailing_regency"`
	MailingDistrict     string `db:"mailing_district"`
	MailingPostalCode   string `db:"mailing_postal_code"`
	MailingFullAddress  string `db:"mailing_full_address"`
	ShippingProvince    string `db:"shipping_province"`
	ShippingRegency     string `db:"shipping_regency"`
	ShippingDistrict    string `db:"shipping_district"`
	ShippingPostalCode  string `db:"shipping_postal_code"`
	ShippingFullAddress string `db:"shipping_full_address"`
}

func (q *Queries) GetContactBooks(ctx context.Context, primaryCompanyID string) ([]GetContactBooksRow, error) {
	rows, err := q.db.QueryContext(ctx, getContactBooks, primaryCompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetContactBooksRow
	for rows.Next() {
		var i GetContactBooksRow
		if err := rows.Scan(
			&i.ID,
			&i.PrimaryCompanyID,
			&i.SecondaryCompanyID,
			&i.ContactGroupID,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Mobile,
			&i.Web,
			&i.IsAllBranches,
			&i.IsCustomer,
			&i.IsSupplier,
			&i.Nickname,
			&i.Tag,
			&i.Note,
			&i.MailingProvince,
			&i.MailingRegency,
			&i.MailingDistrict,
			&i.MailingPostalCode,
			&i.MailingFullAddress,
			&i.ShippingProvince,
			&i.ShippingRegency,
			&i.ShippingDistrict,
			&i.ShippingPostalCode,
			&i.ShippingFullAddress,
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

const getContactGroups = `-- name: GetContactGroups :many
SELECT id, company_id, name
FROM business_relation.contact_groups
WHERE company_id = $1
`

type GetContactGroupsRow struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) GetContactGroups(ctx context.Context, companyID string) ([]GetContactGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getContactGroups, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetContactGroupsRow
	for rows.Next() {
		var i GetContactGroupsRow
		if err := rows.Scan(&i.ID, &i.CompanyID, &i.Name); err != nil {
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

const getCustomers = `-- name: GetCustomers :many
SELECT a.id, a.primary_company_id, a.contact_group_id,
COALESCE(b.name, '') AS contact_group_name, a.name, a.email, a.phone,
a.mobile, a.web, a.is_all_branches, a.is_customer, a.is_supplier,
a.is_tax, a.tax_id, a.is_deleted, COALESCE(c.pic, '') AS pic,
COALESCE(c.credit_limit, 0) AS credit_limit, COALESCE(c.payment_term, 0) AS payment_term
FROM business_relation.contact_books a
LEFT JOIN business_relation.contact_groups b ON a.contact_group_id = b.id
LEFT JOIN business_relation.contact_book_customer_infos c ON a.id = c.contact_book_id
WHERE a.primary_company_id = $1 AND a.is_customer
`

type GetCustomersRow struct {
	ID               string `db:"id"`
	PrimaryCompanyID string `db:"primary_company_id"`
	ContactGroupID   string `db:"contact_group_id"`
	ContactGroupName string `db:"contact_group_name"`
	Name             string `db:"name"`
	Email            string `db:"email"`
	Phone            string `db:"phone"`
	Mobile           string `db:"mobile"`
	Web              string `db:"web"`
	IsAllBranches    bool   `db:"is_all_branches"`
	IsCustomer       bool   `db:"is_customer"`
	IsSupplier       bool   `db:"is_supplier"`
	IsTax            bool   `db:"is_tax"`
	TaxID            string `db:"tax_id"`
	IsDeleted        bool   `db:"is_deleted"`
	Pic              string `db:"pic"`
	CreditLimit      int64  `db:"credit_limit"`
	PaymentTerm      int32  `db:"payment_term"`
}

func (q *Queries) GetCustomers(ctx context.Context, primaryCompanyID string) ([]GetCustomersRow, error) {
	rows, err := q.db.QueryContext(ctx, getCustomers, primaryCompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCustomersRow
	for rows.Next() {
		var i GetCustomersRow
		if err := rows.Scan(
			&i.ID,
			&i.PrimaryCompanyID,
			&i.ContactGroupID,
			&i.ContactGroupName,
			&i.Name,
			&i.Email,
			&i.Phone,
			&i.Mobile,
			&i.Web,
			&i.IsAllBranches,
			&i.IsCustomer,
			&i.IsSupplier,
			&i.IsTax,
			&i.TaxID,
			&i.IsDeleted,
			&i.Pic,
			&i.CreditLimit,
			&i.PaymentTerm,
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

const insertContactBook = `-- name: InsertContactBook :one
INSERT INTO business_relation.contact_books(id, primary_company_id, secondary_company_id,
contact_group_id, name, email, phone, mobile, web,
is_all_branches, is_customer, is_supplier)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,
$10, $11, $12)
RETURNING id, primary_company_id, secondary_company_id, contact_group_id, name, email, phone, mobile, web, is_all_branches, is_customer, is_supplier, is_tax, tax_id, is_deleted, created_at, updated_at
`

type InsertContactBookParams struct {
	ID                 string `db:"id"`
	PrimaryCompanyID   string `db:"primary_company_id"`
	SecondaryCompanyID string `db:"secondary_company_id"`
	ContactGroupID     string `db:"contact_group_id"`
	Name               string `db:"name"`
	Email              string `db:"email"`
	Phone              string `db:"phone"`
	Mobile             string `db:"mobile"`
	Web                string `db:"web"`
	IsAllBranches      bool   `db:"is_all_branches"`
	IsCustomer         bool   `db:"is_customer"`
	IsSupplier         bool   `db:"is_supplier"`
}

func (q *Queries) InsertContactBook(ctx context.Context, arg InsertContactBookParams) (BusinessRelationContactBook, error) {
	row := q.db.QueryRowContext(ctx, insertContactBook,
		arg.ID,
		arg.PrimaryCompanyID,
		arg.SecondaryCompanyID,
		arg.ContactGroupID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Mobile,
		arg.Web,
		arg.IsAllBranches,
		arg.IsCustomer,
		arg.IsSupplier,
	)
	var i BusinessRelationContactBook
	err := row.Scan(
		&i.ID,
		&i.PrimaryCompanyID,
		&i.SecondaryCompanyID,
		&i.ContactGroupID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Mobile,
		&i.Web,
		&i.IsAllBranches,
		&i.IsCustomer,
		&i.IsSupplier,
		&i.IsTax,
		&i.TaxID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertContactBookAdditionalInfo = `-- name: InsertContactBookAdditionalInfo :exec
INSERT INTO business_relation.contact_book_additional_infos(contact_book_id,
nickname, tag, note)
VALUES ($1, $2, $3, $4)
`

type InsertContactBookAdditionalInfoParams struct {
	ContactBookID string `db:"contact_book_id"`
	Nickname      string `db:"nickname"`
	Tag           string `db:"tag"`
	Note          string `db:"note"`
}

func (q *Queries) InsertContactBookAdditionalInfo(ctx context.Context, arg InsertContactBookAdditionalInfoParams) error {
	_, err := q.db.ExecContext(ctx, insertContactBookAdditionalInfo,
		arg.ContactBookID,
		arg.Nickname,
		arg.Tag,
		arg.Note,
	)
	return err
}

const insertContactBookBranch = `-- name: InsertContactBookBranch :exec
INSERT INTO business_relation.contact_book_branches(contact_book_id, company_branch_id)
VALUES ($1, $2)
`

type InsertContactBookBranchParams struct {
	ContactBookID   string `db:"contact_book_id"`
	CompanyBranchID string `db:"company_branch_id"`
}

func (q *Queries) InsertContactBookBranch(ctx context.Context, arg InsertContactBookBranchParams) error {
	_, err := q.db.ExecContext(ctx, insertContactBookBranch, arg.ContactBookID, arg.CompanyBranchID)
	return err
}

const insertContactBookMailingAddress = `-- name: InsertContactBookMailingAddress :exec
INSERT INTO business_relation.contact_book_mailing_addresses(contact_book_id,
province, regency, district, postal_code, full_address)
VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertContactBookMailingAddressParams struct {
	ContactBookID string `db:"contact_book_id"`
	Province      string `db:"province"`
	Regency       string `db:"regency"`
	District      string `db:"district"`
	PostalCode    string `db:"postal_code"`
	FullAddress   string `db:"full_address"`
}

func (q *Queries) InsertContactBookMailingAddress(ctx context.Context, arg InsertContactBookMailingAddressParams) error {
	_, err := q.db.ExecContext(ctx, insertContactBookMailingAddress,
		arg.ContactBookID,
		arg.Province,
		arg.Regency,
		arg.District,
		arg.PostalCode,
		arg.FullAddress,
	)
	return err
}

const insertContactBookShippingAddress = `-- name: InsertContactBookShippingAddress :exec
INSERT INTO business_relation.contact_book_shipping_addresses(contact_book_id,
province, regency, district, postal_code, full_address)
VALUES ($1, $2, $3, $4, $5, $6)
`

type InsertContactBookShippingAddressParams struct {
	ContactBookID string `db:"contact_book_id"`
	Province      string `db:"province"`
	Regency       string `db:"regency"`
	District      string `db:"district"`
	PostalCode    string `db:"postal_code"`
	FullAddress   string `db:"full_address"`
}

func (q *Queries) InsertContactBookShippingAddress(ctx context.Context, arg InsertContactBookShippingAddressParams) error {
	_, err := q.db.ExecContext(ctx, insertContactBookShippingAddress,
		arg.ContactBookID,
		arg.Province,
		arg.Regency,
		arg.District,
		arg.PostalCode,
		arg.FullAddress,
	)
	return err
}

const insertContactGroup = `-- name: InsertContactGroup :one
INSERT INTO business_relation.contact_groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING id, company_id, name, created_at, updated_at
`

type InsertContactGroupParams struct {
	ID        string `db:"id"`
	CompanyID string `db:"company_id"`
	Name      string `db:"name"`
}

func (q *Queries) InsertContactGroup(ctx context.Context, arg InsertContactGroupParams) (BusinessRelationContactGroup, error) {
	row := q.db.QueryRowContext(ctx, insertContactGroup, arg.ID, arg.CompanyID, arg.Name)
	var i BusinessRelationContactGroup
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateContactBook = `-- name: UpdateContactBook :one
UPDATE business_relation.contact_books
SET 
    contact_group_id = $2,
    name = $3,
    email = $4,
    phone = $5,
    mobile = $6,
    web = $7,
    is_all_branches = $8,
    is_customer = $9,
    is_supplier = $10,
    updated_at = NOW()
WHERE id = $1
RETURNING id, primary_company_id, secondary_company_id, contact_group_id, name, email, phone, mobile, web, is_all_branches, is_customer, is_supplier, is_tax, tax_id, is_deleted, created_at, updated_at
`

type UpdateContactBookParams struct {
	ID             string `db:"id"`
	ContactGroupID string `db:"contact_group_id"`
	Name           string `db:"name"`
	Email          string `db:"email"`
	Phone          string `db:"phone"`
	Mobile         string `db:"mobile"`
	Web            string `db:"web"`
	IsAllBranches  bool   `db:"is_all_branches"`
	IsCustomer     bool   `db:"is_customer"`
	IsSupplier     bool   `db:"is_supplier"`
}

func (q *Queries) UpdateContactBook(ctx context.Context, arg UpdateContactBookParams) (BusinessRelationContactBook, error) {
	row := q.db.QueryRowContext(ctx, updateContactBook,
		arg.ID,
		arg.ContactGroupID,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.Mobile,
		arg.Web,
		arg.IsAllBranches,
		arg.IsCustomer,
		arg.IsSupplier,
	)
	var i BusinessRelationContactBook
	err := row.Scan(
		&i.ID,
		&i.PrimaryCompanyID,
		&i.SecondaryCompanyID,
		&i.ContactGroupID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.Mobile,
		&i.Web,
		&i.IsAllBranches,
		&i.IsCustomer,
		&i.IsSupplier,
		&i.IsTax,
		&i.TaxID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateContactBookAdditionalInfo = `-- name: UpdateContactBookAdditionalInfo :exec
UPDATE business_relation.contact_book_additional_infos
SET 
    nickname = $2,
    tag = $3,
    note = $4,
    updated_at = NOW()
WHERE contact_book_id = $1
`

type UpdateContactBookAdditionalInfoParams struct {
	ContactBookID string `db:"contact_book_id"`
	Nickname      string `db:"nickname"`
	Tag           string `db:"tag"`
	Note          string `db:"note"`
}

func (q *Queries) UpdateContactBookAdditionalInfo(ctx context.Context, arg UpdateContactBookAdditionalInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateContactBookAdditionalInfo,
		arg.ContactBookID,
		arg.Nickname,
		arg.Tag,
		arg.Note,
	)
	return err
}

const updateContactBookMailingAddress = `-- name: UpdateContactBookMailingAddress :exec
UPDATE business_relation.contact_book_mailing_addresses
SET 
    province = $2,
    regency = $3,
    district = $4,
    postal_code = $5,
    full_address = $6,
    updated_at = NOW()
WHERE contact_book_id = $1
`

type UpdateContactBookMailingAddressParams struct {
	ContactBookID string `db:"contact_book_id"`
	Province      string `db:"province"`
	Regency       string `db:"regency"`
	District      string `db:"district"`
	PostalCode    string `db:"postal_code"`
	FullAddress   string `db:"full_address"`
}

func (q *Queries) UpdateContactBookMailingAddress(ctx context.Context, arg UpdateContactBookMailingAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateContactBookMailingAddress,
		arg.ContactBookID,
		arg.Province,
		arg.Regency,
		arg.District,
		arg.PostalCode,
		arg.FullAddress,
	)
	return err
}

const updateContactBookShippingAddress = `-- name: UpdateContactBookShippingAddress :exec
UPDATE business_relation.contact_book_shipping_addresses
SET 
    province = $2,
    regency = $3,
    district = $4,
    postal_code = $5,
    full_address = $6,
    updated_at = NOW()
WHERE contact_book_id = $1
`

type UpdateContactBookShippingAddressParams struct {
	ContactBookID string `db:"contact_book_id"`
	Province      string `db:"province"`
	Regency       string `db:"regency"`
	District      string `db:"district"`
	PostalCode    string `db:"postal_code"`
	FullAddress   string `db:"full_address"`
}

func (q *Queries) UpdateContactBookShippingAddress(ctx context.Context, arg UpdateContactBookShippingAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateContactBookShippingAddress,
		arg.ContactBookID,
		arg.Province,
		arg.Regency,
		arg.District,
		arg.PostalCode,
		arg.FullAddress,
	)
	return err
}

const updateContactBookTaxInfo = `-- name: UpdateContactBookTaxInfo :exec
UPDATE business_relation.contact_books
SET 
    is_tax = $2,
    tax_id = $3,
    updated_at = NOW()
WHERE id = $1
`

type UpdateContactBookTaxInfoParams struct {
	ID    string `db:"id"`
	IsTax bool   `db:"is_tax"`
	TaxID string `db:"tax_id"`
}

func (q *Queries) UpdateContactBookTaxInfo(ctx context.Context, arg UpdateContactBookTaxInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateContactBookTaxInfo, arg.ID, arg.IsTax, arg.TaxID)
	return err
}

const updateContactGroup = `-- name: UpdateContactGroup :one
UPDATE business_relation.contact_groups
SET 
    name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, company_id, name, created_at, updated_at
`

type UpdateContactGroupParams struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func (q *Queries) UpdateContactGroup(ctx context.Context, arg UpdateContactGroupParams) (BusinessRelationContactGroup, error) {
	row := q.db.QueryRowContext(ctx, updateContactGroup, arg.ID, arg.Name)
	var i BusinessRelationContactGroup
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertCustomerInfo = `-- name: UpsertCustomerInfo :exec
INSERT INTO business_relation.contact_book_customer_infos(contact_book_id, pic, credit_limit, payment_term)
VALUES ($1, $2, $3, $4)
ON CONFLICT (contact_book_id)
DO UPDATE SET
    pic = EXCLUDED.pic,
    credit_limit = EXCLUDED.credit_limit,
    payment_term = EXCLUDED.payment_term,
    updated_at = NOW()
`

type UpsertCustomerInfoParams struct {
	ContactBookID string `db:"contact_book_id"`
	Pic           string `db:"pic"`
	CreditLimit   int64  `db:"credit_limit"`
	PaymentTerm   int32  `db:"payment_term"`
}

func (q *Queries) UpsertCustomerInfo(ctx context.Context, arg UpsertCustomerInfoParams) error {
	_, err := q.db.ExecContext(ctx, upsertCustomerInfo,
		arg.ContactBookID,
		arg.Pic,
		arg.CreditLimit,
		arg.PaymentTerm,
	)
	return err
}

const upsertSupplierInfo = `-- name: UpsertSupplierInfo :exec
INSERT INTO business_relation.contact_book_supplier_infos(contact_book_id, pic, credit_limit, payment_term)
VALUES ($1, $2, $3, $4)
ON CONFLICT (contact_book_id)
DO UPDATE SET
    pic = EXCLUDED.pic,
    credit_limit = EXCLUDED.credit_limit,
    payment_term = EXCLUDED.payment_term,
    updated_at = NOW()
`

type UpsertSupplierInfoParams struct {
	ContactBookID string `db:"contact_book_id"`
	Pic           string `db:"pic"`
	CreditLimit   int64  `db:"credit_limit"`
	PaymentTerm   int32  `db:"payment_term"`
}

func (q *Queries) UpsertSupplierInfo(ctx context.Context, arg UpsertSupplierInfoParams) error {
	_, err := q.db.ExecContext(ctx, upsertSupplierInfo,
		arg.ContactBookID,
		arg.Pic,
		arg.CreditLimit,
		arg.PaymentTerm,
	)
	return err
}
