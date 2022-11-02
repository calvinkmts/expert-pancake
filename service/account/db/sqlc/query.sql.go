// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO account.users (id, fullname, nickname, email, phone_number)
VALUES ($1, $2, $3, $4, $5)
`

type CreateUserParams struct {
	ID          string         `db:"id"`
	Fullname    string         `db:"fullname"`
	Nickname    string         `db:"nickname"`
	Email       sql.NullString `db:"email"`
	PhoneNumber string         `db:"phone_number"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Fullname,
		arg.Nickname,
		arg.Email,
		arg.PhoneNumber,
	)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, fullname, nickname, email, phone_number, created_at, updated_at FROM account.users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (AccountUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i AccountUser
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Nickname,
		&i.Email,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByPhoneNumber = `-- name: GetUserByPhoneNumber :one
SELECT id FROM account.users
WHERE phone_number = $1
`

func (q *Queries) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhoneNumber, phoneNumber)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getUserPassword = `-- name: GetUserPassword :one
SELECT user_id, password, created_at, updated_at FROM account.user_passwords
WHERE user_id = $1
`

func (q *Queries) GetUserPassword(ctx context.Context, userID string) (AccountUserPassword, error) {
	row := q.db.QueryRowContext(ctx, getUserPassword, userID)
	var i AccountUserPassword
	err := row.Scan(
		&i.UserID,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const upsertUserAddresses = `-- name: UpsertUserAddresses :exec
INSERT INTO account.user_addresses(user_id, country, province, regency, district, full_address)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (user_id)
DO UPDATE SET
    country = EXCLUDED.country,
    province = EXCLUDED.province,
    regency = EXCLUDED.regency,
    district = EXCLUDED.district,
    full_address = EXCLUDED.full_address,
    updated_at = NOW()
`

type UpsertUserAddressesParams struct {
	UserID      string `db:"user_id"`
	Country     string `db:"country"`
	Province    string `db:"province"`
	Regency     string `db:"regency"`
	District    string `db:"district"`
	FullAddress string `db:"full_address"`
}

func (q *Queries) UpsertUserAddresses(ctx context.Context, arg UpsertUserAddressesParams) error {
	_, err := q.db.ExecContext(ctx, upsertUserAddresses,
		arg.UserID,
		arg.Country,
		arg.Province,
		arg.Regency,
		arg.District,
		arg.FullAddress,
	)
	return err
}

const upsertUserInfo = `-- name: UpsertUserInfo :exec
INSERT INTO account.user_infos(user_id, key, value)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, key)
DO UPDATE SET
    value = EXCLUDED.value,
    updated_at = NOW()
`

type UpsertUserInfoParams struct {
	UserID string `db:"user_id"`
	Key    string `db:"key"`
	Value  string `db:"value"`
}

func (q *Queries) UpsertUserInfo(ctx context.Context, arg UpsertUserInfoParams) error {
	_, err := q.db.ExecContext(ctx, upsertUserInfo, arg.UserID, arg.Key, arg.Value)
	return err
}

const upsertUserPassword = `-- name: UpsertUserPassword :exec
INSERT INTO account.user_passwords(user_id, password)
VALUES ($1, $2)
ON CONFLICT (user_id)
DO UPDATE SET
    password = EXCLUDED.password,
    updated_at = NOW()
`

type UpsertUserPasswordParams struct {
	UserID   string `db:"user_id"`
	Password string `db:"password"`
}

func (q *Queries) UpsertUserPassword(ctx context.Context, arg UpsertUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, upsertUserPassword, arg.UserID, arg.Password)
	return err
}
