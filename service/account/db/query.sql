-- name: CreateUser :exec
INSERT INTO account.users (id, image_url, fullname, nickname, email, phone_number)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpsertUserInfo :exec
INSERT INTO account.user_infos(user_id, key, value)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, key)
DO UPDATE SET
    value = EXCLUDED.value,
    updated_at = NOW();

-- name: UpsertUserPassword :exec
INSERT INTO account.user_passwords(user_id, password)
VALUES ($1, $2)
ON CONFLICT (user_id)
DO UPDATE SET
    password = EXCLUDED.password,
    updated_at = NOW();

-- name: GetUserPassword :one
SELECT * FROM account.user_passwords
WHERE user_id = $1;

-- name: GetUser :one
SELECT * FROM account.users
WHERE id = $1;

-- name: GetUserByPhoneNumber :one
SELECT id FROM account.users
WHERE phone_number = $1;

-- name: UpsertUserAddresses :one
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
RETURNING *;

-- name: UpdateUser :one
UPDATE account.users
SET
    image_url = $2,
    fullname = $3,
    nickname = $4,
    email = $5,
    phone_number = $6,
    updated_at = NOW()
WHERE
    id = $1
RETURNING *;

-- name: GetUserInfo :one
SELECT * FROM account.user_infos
WHERE user_id = $1 AND key = $2;

-- name: UpsertUser :one
INSERT INTO account.users (id, image_url, fullname, nickname, email, phone_number)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id)
DO UPDATE SET
    image_url = EXCLUDED.image_url,
    fullname = EXCLUDED.fullname,
    nickname = EXCLUDED.nickname,
    email = EXCLUDED.email,
    phone_number = EXCLUDED.phone_number
RETURNING *;

-- name: GetUserAddress :one
SELECT * FROM account.user_addresses
WHERE user_id = $1;