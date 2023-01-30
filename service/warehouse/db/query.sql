-- name: UpsertRack :one
INSERT INTO warehouse.racks(id, branch_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    branch_id = EXCLUDED.branch_id,
    updated_at = NOW()
RETURNING *;

-- name: GetRacks :many
SELECT a.id, a.branch_id, a.name
FROM warehouse.racks a
LEFT JOIN warehouse.warehouse_racks b ON a.id = b.rack_id
WHERE a.branch_id = $1 
    AND a.name LIKE $2 
    AND CASE WHEN
        @IsGetAvailable::bool THEN b.rack_id IS NULL ELSE TRUE END;

-- name: GetWarehouses :many
SELECT a.*
FROM warehouse.warehouses a
WHERE  a.branch_id = $1
AND a.name LIKE $2
AND a.type LIKE $3
AND a.is_deleted = FALSE;

-- name: AddWarehouse :one
INSERT INTO warehouse.warehouses
(id, branch_id, code, name, address, type)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: AddWarehouseRack :one
INSERT INTO warehouse.warehouse_racks
(warehouse_id, rack_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetRack :one
SELECT id, branch_id, name
FROM warehouse.racks 
WHERE id = $1;

-- name: GetWarehouseRacks :many
SELECT b.id, b.branch_id, b.name
FROM warehouse.warehouse_racks a JOIN warehouse.racks b ON a.rack_id = b.id 
WHERE a.warehouse_id = $1;
