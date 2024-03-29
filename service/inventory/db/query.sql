-- name: InsertBrand :one
INSERT INTO inventory.brands(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBrand :one
UPDATE inventory.brands
SET name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteBrand :exec
UPDATE inventory.brands
SET is_deleted = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- name: GetBrands :many
SELECT id,
    company_id,
    name
FROM inventory.brands
WHERE company_id = $1
    AND name LIKE $2
    AND is_deleted = false;

-- name: InsertGroup :one
INSERT INTO inventory.groups(id, company_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateGroup :one
UPDATE inventory.groups
SET name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
UPDATE inventory.groups
SET is_deleted = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- name: GetGroups :many
SELECT id,
    company_id,
    name
FROM inventory.groups
WHERE company_id = $1
    AND name LIKE $2
    AND is_deleted = false;

-- name: InsertUnit :one
INSERT INTO inventory.units(id, company_id, unit_category_id, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUnit :one
UPDATE inventory.units
SET unit_category_id = $2,
    name = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetUnits :many
SELECT id,
    company_id,
    unit_category_id,
    name
FROM inventory.units
WHERE company_id = $1
    AND unit_category_id LIKE $2
    AND name LIKE $3;

-- name: InsertItem :one
INSERT INTO inventory.items(
        id,
        company_id,
        image_url,
        code,
        name,
        brand_id,
        group_id,
        tag,
        description
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: InsertItemVariant :one
INSERT INTO inventory.item_variants(
        id,
        item_id,
        image_url,
        barcode,
        name,
        price,
        is_default
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteItemVariant :exec
UPDATE inventory.item_variants
SET is_deleted = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- name: GetBrandById :one
SELECT id,
    company_id,
    name
FROM inventory.brands
WHERE id = $1;

-- name: GetGroupById :one
SELECT id,
    company_id,
    name
FROM inventory.groups
WHERE id = $1;

-- name: GetItemGroups :one
SELECT string_agg(CONCAT(id, '|', name), ',')::text AS groups
FROM inventory.groups
WHERE is_deleted = false
AND id = ANY(@group_ids::text []);

-- name: UpdateItem :one
UPDATE inventory.items
SET image_url = $2,
    name = $3,
    brand_id = $4,
    group_id = $5,
    tag = $6,
    description = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateItemVariantDefault :one
UPDATE inventory.item_variants
SET image_url = $2,
    barcode = $3,
    price = $4,
    updated_at = NOW()
WHERE item_id = $1
    AND is_default = true
RETURNING *;

-- name: GetItems :many
SELECT a.id,
    b.id AS variant_id,
    a.company_id,
    b.image_url,
    a.code,
    b.barcode,
    a.name,
    b.name AS variant_name,
    a.brand_id,
    COALESCE(c.name, '') AS brand_name,
    a.group_id,
    string_agg(CONCAT(d.id, '|', d.name), ',')::text AS groups,
    COALESCE(SUM(rp.amount), 0)::bigint AS amount_in_stock,
    a.tag,
    a.description,
    b.is_default,
    b.price
FROM inventory.items a
    JOIN inventory.item_variants b ON a.id = b.item_id
    LEFT JOIN inventory.brands c ON a.brand_id = c.id
    LEFT JOIN inventory.stock_movements rp ON rp.variant_id = b.id
    JOIN inventory.groups d ON d.id = ANY(string_to_array(a.group_id, ','))
WHERE a.company_id = $1
    AND (b.name LIKE @keyword OR a.tag LIKE @keyword)
GROUP BY a.id, b.id, c.id;

-- name: UpsertItemInfo :exec
INSERT INTO inventory.item_infos(
        item_id,
        is_purchase,
        is_sale,
        is_raw_material,
        is_asset,
        purchase_chart_of_account_id,
        sale_chart_of_account_id,
        purchase_item_unit_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ON CONFLICT (item_id) DO
UPDATE
SET is_purchase = EXCLUDED.is_purchase,
    is_sale = EXCLUDED.is_sale,
    is_raw_material = EXCLUDED.is_raw_material,
    is_asset = EXCLUDED.is_asset,
    purchase_chart_of_account_id = EXCLUDED.purchase_chart_of_account_id,
    sale_chart_of_account_id = EXCLUDED.sale_chart_of_account_id,
    purchase_item_unit_id = EXCLUDED.purchase_item_unit_id,
    updated_at = NOW();

-- name: GetItemInfo :one
SELECT a.item_id,
    d.company_id,
    a.is_purchase,
    a.is_sale,
    a.is_raw_material,
    a.is_asset,
    a.purchase_chart_of_account_id,
    a.sale_chart_of_account_id,
    a.purchase_item_unit_id,
    c.name AS purchase_item_unit_name
FROM inventory.item_infos a
    JOIN inventory.item_units b ON a.purchase_item_unit_id = b.id
    JOIN inventory.units c ON b.unit_id = c.id
    JOIN inventory.items d ON a.item_id = d.id
WHERE a.item_id = $1;

-- name: UpsertItemVariant :exec
INSERT INTO inventory.item_variants(id, item_id, image_url, barcode, name, price)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO
UPDATE
SET item_id = EXCLUDED.item_id,
    image_url = EXCLUDED.image_url,
    barcode = EXCLUDED.barcode,
    name = EXCLUDED.name,
    price = EXCLUDED.price,
    updated_at = NOW();

-- name: GetItemVariant :one
SELECT b.id,
    a.id AS variant_id,
    b.company_id,
    a.image_url,
    b.code,
    a.barcode,
    b.name,
    a.name AS variant_name,
    b.brand_id,
    COALESCE(c.name, '') AS brand_name,
    b.group_id,
    string_agg(CONCAT(d.id, '|', d.name), ',')::text AS groups,
    b.tag,
    b.description,
    a.is_default,
    a.price
FROM inventory.item_variants a
    JOIN inventory.items b ON a.item_id = b.id
    LEFT JOIN inventory.brands c ON b.brand_id = c.id
    JOIN inventory.groups d ON d.id = ANY(string_to_array(b.group_id, ','))
WHERE a.id = $1
GROUP BY a.id, b.id, c.id;

-- name: GetItemVariants :many
SELECT b.id,
    a.id AS variant_id,
    b.company_id,
    a.image_url,
    b.code,
    a.barcode,
    b.name,
    a.name AS variant_name,
    b.brand_id,
    COALESCE(c.name, '') AS brand_name,
    b.group_id,
    string_agg(CONCAT(d.id, '|', d.name), ',')::text AS groups,
    b.tag,
    b.description,
    a.is_default,
    a.price
FROM inventory.item_variants a
    JOIN inventory.items b ON a.item_id = b.id
    LEFT JOIN inventory.brands c ON b.brand_id = c.id
    JOIN inventory.groups d ON d.id = ANY(string_to_array(b.group_id, ','))
WHERE a.is_deleted = FALSE AND CASE WHEN @is_filter_id::bool THEN a.id = $1 ELSE a.item_id = $2
    AND a.name LIKE $3 END
GROUP BY a.id, b.id, c.id;

-- name: UpsertItemUnit :one
INSERT INTO inventory.item_units(id, item_id, unit_id, value, is_default)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO
UPDATE
SET item_id = EXCLUDED.item_id,
    unit_id = EXCLUDED.unit_id,
    value = EXCLUDED.value,
    is_default = EXCLUDED.is_default,
    updated_at = NOW()
RETURNING *;

-- name: DeleteItemUnit :exec
UPDATE inventory.item_units
SET is_deleted = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateItemUnitIsDefaultToFalse :exec
UPDATE inventory.item_units
SET is_default = false,
    updated_at = NOW()
WHERE item_id = $1
AND id <> $2;

-- name: GetUnit :one
SELECT id,
    company_id,
    name
FROM inventory.units
WHERE id = $1;

-- name: GetItemUnits :many
SELECT a.id,
    a.item_id,
    a.unit_id,
    b.name AS unit_name,
    a.value,
    a.is_default
FROM inventory.item_units a
    JOIN inventory.units b ON a.unit_id = b.id
WHERE CASE WHEN @is_filter_id::bool THEN a.id = $1 ELSE a.item_id = $2
    AND b.name LIKE $3
	AND a.is_deleted = false END;

-- name: InsertInternalStockTransfer :one
INSERT INTO inventory.internal_stock_transfers(
        id,
        company_id,
        branch_id,
        source_warehouse_id,
        destination_warehouse_id,
        form_number,
        transaction_date
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetInternalStockTransfers :many
SELECT id,
    source_warehouse_id,
    destination_warehouse_id,
    form_number,
    transaction_date
FROM inventory.internal_stock_transfers
WHERE is_deleted = false
    AND transaction_date BETWEEN @start_date::date AND @end_date::date
    AND (
        source_warehouse_id = ANY(@warehouse_ids::text [])
        OR destination_warehouse_id = ANY(@warehouse_ids::text [])
    );

-- name: InsertInternalStockTransferItem :exec
INSERT INTO inventory.internal_stock_transfer_items(
        id,
        internal_stock_transfer_id,
        warehouse_rack_id,
        variant_id,
        item_unit_id,
        item_unit_value,
        amount,
        batch,
        expired_date,
        item_barcode_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetInternalStockTransferItems :many
SELECT a.id,
    a.warehouse_rack_id,
    e.name AS item_name,
    a.variant_id,
    b.name AS variant_name,
    a.item_unit_id,
    d.name AS item_unit_name,
    a.item_unit_value,
    a.amount,
    a.batch,
    a.expired_date
FROM inventory.internal_stock_transfer_items a
    JOIN inventory.item_variants b ON a.variant_id = b.id
    JOIN inventory.item_units c ON a.item_unit_id = c.id
    JOIN inventory.units d ON c.unit_id = d.id
    JOIN inventory.items e ON b.item_id = e.id
WHERE a.internal_stock_transfer_id = $1
    AND a.is_deleted = false;

-- name: InsertItemBarcode :exec
INSERT INTO inventory.item_barcodes(id, variant_id, batch, expired_date)
VALUES ($1, $2, $3, $4);

-- name: GetItemBarcode :one
SELECT id
FROM inventory.item_barcodes
WHERE variant_id = $1
    AND CASE
        WHEN @is_null_batch::bool THEN batch is null
        ELSE batch = $2
    END
    AND CASE
        WHEN @is_null_expired_date::bool THEN expired_date is null
        ELSE expired_date = $3
    END;

-- name: InsertStockMovement :exec
INSERT INTO inventory.stock_movements(
        id,
        transaction_id,
        company_id,
        branch_id,
        transaction_code,
        transaction_date,
        transaction_reference,
        detail_transaction_id,
        warehouse_id,
        warehouse_rack_id,
        variant_id,
        item_barcode_id,
        amount
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);

-- name: DeleteStockMovement :exec
DELETE FROM inventory.stock_movements
WHERE transaction_id = $1 AND transaction_reference = $2;

-- name: InsertUpdateStock :exec
INSERT INTO inventory.update_stocks(id,
form_number, transaction_date, company_id, branch_id, warehouse_id, warehouse_rack_id,
variant_id, item_unit_id, item_unit_value, beginning_stock, ending_stock,
batch, expired_date, item_barcode_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);

-- name: GetUpdateStock :one
SELECT a.id, a.form_number, a.transaction_date, a.warehouse_id, a.warehouse_rack_id,
b.item_id, c.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, e.name AS item_unit_name, a.item_unit_value,
a.beginning_stock, a.ending_stock, a.batch, a.expired_date
FROM inventory.update_stocks a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.item_units d ON a.item_unit_id = d.id
JOIN inventory.units e ON d.unit_id = e.id
WHERE a.id = $1;

-- name: GetUpdateStocks :many
SELECT a.id, a.form_number, a.transaction_date, a.warehouse_id, a.warehouse_rack_id,
b.item_id, c.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, e.name AS item_unit_name, a.item_unit_value,
a.beginning_stock, a.ending_stock, a.batch, a.expired_date
FROM inventory.update_stocks a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.item_units d ON a.item_unit_id = d.id
JOIN inventory.units e ON d.unit_id = e.id
WHERE a.is_deleted = false
AND a.transaction_date BETWEEN @start_date::date AND @end_date::date
AND a.warehouse_id = ANY(@warehouse_ids::text[]);

-- name: UpsertItemReorder :one
INSERT INTO inventory.item_reorders(
        id,
        variant_id,
        item_unit_id,
        warehouse_id,
        minimum_stock
    )
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (id) DO
UPDATE
SET variant_id = EXCLUDED.variant_id,
    item_unit_id = EXCLUDED.item_unit_id,
    warehouse_id = EXCLUDED.warehouse_id,
    minimum_stock = EXCLUDED.minimum_stock,
    updated_at = NOW()
RETURNING *;

-- name: GetItemReorder :one
SELECT a.id,
    a.variant_id,
    b.name as variant_name,
    c.id as item_id,
    c.name as item_name,
    d.id as item_unit_id,
    e.name as item_unit_name,
    a.warehouse_id,
    a.minimum_stock
FROM inventory.item_reorders a
    JOIN inventory.item_variants b ON a.variant_id = b.id
    JOIN inventory.items c ON b.item_id = c.id
    JOIN inventory.item_units d ON a.item_unit_id = d.id
    JOIN inventory.units e ON d.unit_id = e.id
WHERE a.id = $1;

-- name: GetItemReorders :many
SELECT a.id,
    a.variant_id,
    b.name as variant_name,
    c.id as item_id,
    c.name as item_name,
    d.id as item_unit_id,
    e.name as item_unit_name,
    a.warehouse_id,
    a.minimum_stock
FROM inventory.item_reorders a
    JOIN inventory.item_variants b ON a.variant_id = b.id
    JOIN inventory.items c ON b.item_id = c.id
    JOIN inventory.item_units d ON a.item_unit_id = d.id
    JOIN inventory.units e ON d.unit_id = e.id
WHERE a.warehouse_id LIKE $1
    AND b.item_id LIKE $2;

-- name: UpsertUnitCategory :one
INSERT INTO inventory.unit_categories(id, company_id, name)
VALUES ($1, $2, $3) ON CONFLICT (id) DO
UPDATE
SET company_id = EXCLUDED.company_id,
    name = EXCLUDED.name,
    updated_at = NOW()
RETURNING *;

-- name: GetUnitCategories :many
SELECT id,
    company_id,
    name
FROM inventory.unit_categories
WHERE company_id = $1
    AND name LIKE $2;

-- name: GetVariantWarehouseRacks :many
SELECT DISTINCT a.warehouse_rack_id
FROM inventory.stock_movements a
WHERE a.variant_id = $1
AND a.warehouse_id = $2;

-- name: GetVariantWarehouseRackBatches :many
SELECT DISTINCT b.batch
FROM inventory.stock_movements a
JOIN inventory.item_barcodes b ON a.item_barcode_id = b.id
WHERE a.variant_id = $1
AND a.warehouse_rack_id = $2;

-- name: GetVariantWarehouseRackBatchExpiredDates :many
SELECT DISTINCT b.expired_date
FROM inventory.stock_movements a
JOIN inventory.item_barcodes b ON a.item_barcode_id = b.id
WHERE a.variant_id = $1
AND a.warehouse_rack_id = $2
AND CASE
    WHEN @is_null_batch::bool THEN b.batch is null
    ELSE b.batch = $3
END;

-- name: GetVariantWarehouseRackStock :one
SELECT b.id, COALESCE(SUM(a.amount), 0)::bigint AS stock
FROM inventory.stock_movements a
JOIN inventory.item_barcodes b ON a.item_barcode_id = b.id
WHERE a.variant_id = $1
AND a.warehouse_rack_id = $2
AND CASE
    WHEN @is_null_batch::bool THEN b.batch is null
    ELSE b.batch = $3
END
AND CASE
    WHEN @is_null_expired_date::bool THEN b.expired_date is null
    ELSE b.expired_date = $4
END
GROUP BY b.id;

-- name: GetVariantWarehouseStocks :many
SELECT c.id AS item_id, c.name AS item_name,
b.id AS variant_id, b.name AS variant_name,
SUM(a.amount) AS stock
FROM inventory.stock_movements a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
WHERE a.warehouse_id = $1
GROUP BY c.id, b.id;

-- name: GetTransferHistory :many
SELECT b.transaction_date,
    b.form_number,
    item.id as item_id,
    item.name as item_name,
    item.image_url,
    variant.id as variant_id,
    variant.name as variant_name,
    b.source_warehouse_id,
    b.destination_warehouse_id,
    a.amount
FROM inventory.internal_stock_transfer_items a
    JOIN inventory.item_variants variant ON variant.id = a.variant_id
    JOIN inventory.items item ON item.id = variant.item_id
    JOIN inventory.internal_stock_transfers b ON b.id = a.internal_stock_transfer_id
WHERE a.is_deleted = false
    AND b.is_deleted = false
    AND variant.item_id LIKE $1
    AND transaction_date BETWEEN @start_date::date AND @end_date::date
    AND CASE WHEN @is_received_filter::bool
        THEN b.is_received = $2 ELSE TRUE END
    AND (
        source_warehouse_id LIKE $3
        OR destination_warehouse_id LIKE $4
    )
    AND (
        source_warehouse_id = ANY(@warehouse_ids::text [])
        OR destination_warehouse_id = ANY(@warehouse_ids::text [])
    )
ORDER BY b.transaction_date DESC;

-- name: GetStockHistory :many
SELECT a.transaction_date,
    a.form_number,
    item.id as item_id,
    item.name as item_name,
    item.image_url,
    variant.id as variant_id,
    variant.name as variant_name,
    a.beginning_stock as onhand,
    a.ending_stock as calculated
FROM inventory.update_stocks a
    JOIN inventory.item_variants variant ON variant.id = a.variant_id
    JOIN inventory.items item ON item.id = variant.item_id
WHERE a.is_deleted = false
    AND variant.item_id LIKE $1
    AND transaction_date BETWEEN @start_date::date AND @end_date::date
    AND warehouse_id = ANY(@warehouse_ids::text [])
ORDER BY a.transaction_date DESC;

-- name: GetItemReorderNotifications :many
SELECT b.id AS item_id, b.name AS item_name,
a.id AS variant_id, a.name AS variant_name,
COALESCE(SUM(c.amount), 0)::bigint AS current_stock,
d.id AS item_reorder_id,
COALESCE((d.minimum_stock*e.value), 0)::bigint AS minimum_stock
FROM inventory.item_variants a
JOIN inventory.items b ON a.item_id = b.id
LEFT JOIN inventory.stock_movements c ON a.id = c.variant_id
AND c.warehouse_id = $1
LEFT JOIN inventory.item_reorders d ON a.id = d.variant_id
AND d.warehouse_id = $1
LEFT JOIN inventory.item_units e ON b.id = e.item_id
AND d.item_unit_id = e.id
WHERE a.id = ANY(@item_variant_ids::text [])
GROUP BY b.id, a.id, d.id, e.value;

-- name: GetSupplierCatalogs :many
SELECT a.id,
    b.id AS variant_id,
    a.company_id,
    b.image_url,
    a.code,
    b.barcode,
    a.name,
    b.name AS variant_name,
    a.brand_id,
    COALESCE(c.name, '') AS brand_name,
    b.is_default,
    b.price,
    CASE WHEN COUNT(d.id) > 0 THEN true ELSE false END AS is_mapped
FROM inventory.items a
    JOIN inventory.item_variants b ON a.id = b.item_id
    LEFT JOIN inventory.brands c ON a.brand_id = c.id
    LEFT JOIN inventory.item_variant_maps d ON b.id = d.primary_item_variant_id
    AND d.secondary_company_id = $2
WHERE a.company_id = $1
    AND b.name LIKE @keyword
    AND CASE WHEN b.is_default THEN
    NOT EXISTS (
		SELECT
			1
		FROM
			inventory.item_variants a1
		WHERE
			a1.item_id = b.item_id
            AND a1.is_default is false
	)
    ELSE TRUE END
GROUP BY a.id, b.id, c.id;

-- name: GetMappingItems :many
SELECT a.id,
    a.code,
    a.name
FROM inventory.items a
WHERE a.company_id = @secondary_company_id
AND a.name LIKE $1
AND NOT EXISTS (
    SELECT
        1
    FROM
        inventory.item_variant_maps a1
    JOIN inventory.item_variants b1 ON a1.secondary_item_variant_id = b1.id
    WHERE
        a1.primary_company_id = @primary_company_id
        AND a1.secondary_company_id = @secondary_company_id
        AND b1.item_id = a.id
)
UNION ALL
SELECT a.id,
    a.code,
    a.name
FROM inventory.items a
JOIN inventory.item_variants b ON a.id = b.item_id
JOIN inventory.item_variant_maps c ON b.id = c.secondary_item_variant_id
AND c.primary_company_id = @primary_company_id
JOIN inventory.item_variants d ON c.primary_item_variant_id = d.id
AND d.item_id = $2
WHERE a.company_id = @secondary_company_id
AND a.name LIKE $1
AND EXISTS (
    SELECT
        1
    FROM
        inventory.item_variants a2
    LEFT JOIN inventory.item_variant_maps b2 ON a2.id = b2.secondary_item_variant_id
    AND b2.primary_company_id = @primary_company_id
    WHERE
        a2.item_id = a.id
        AND b2.id is null
);

-- name: GetMappingItemVariants :many
SELECT b.id,
    a.id AS variant_id,
    b.company_id,
    a.image_url,
    b.code,
    a.barcode,
    b.name,
    a.name AS variant_name,
    b.brand_id,
    COALESCE(c.name, '') AS brand_name,
    a.is_default,
    a.price
FROM inventory.item_variants a
    JOIN inventory.items b ON a.item_id = b.id
    LEFT JOIN inventory.brands c ON b.brand_id = c.id
    LEFT JOIN inventory.item_variant_maps d ON a.id = d.secondary_item_variant_id
    AND d.primary_company_id = $3 AND d.secondary_company_id = $4
WHERE a.item_id = $1
    AND a.name LIKE $2
    AND d.id is null
    AND CASE WHEN a.is_default THEN
    NOT EXISTS (
		SELECT
			1
		FROM
			inventory.item_variants a1
		WHERE
			a1.item_id = a.item_id
            AND a1.is_default is false
	)
    ELSE TRUE END
GROUP BY a.id, b.id, c.id;

-- name: GetMappingItemUnits :many
SELECT c.id AS item_unit_id, d.name AS unit_name
FROM inventory.item_variants a
JOIN inventory.items b ON a.item_id = b.id
JOIN inventory.item_units c ON b.id = c.item_id
JOIN inventory.units d ON c.unit_id = d.id
LEFT JOIN inventory.item_variant_maps e ON a.id = 
CASE WHEN @is_supplier::bool THEN e.primary_item_variant_id ELSE e.secondary_item_variant_id END
AND c.id =
CASE WHEN @is_supplier::bool THEN e.primary_item_unit_id ELSE e.secondary_item_unit_id END
AND CASE WHEN @is_supplier::bool THEN e.secondary_company_id = $2 ELSE e.primary_company_id = $3 END
WHERE a.id = $1
AND e.id is null;

-- name: UpsertItemVariantMap :exec
INSERT INTO inventory.item_variant_maps(id,
primary_company_id, secondary_company_id,
primary_item_variant_id, secondary_item_variant_id,
primary_item_unit_id, secondary_item_unit_id)
VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO
UPDATE
SET primary_company_id = EXCLUDED.primary_company_id,
    secondary_company_id = EXCLUDED.secondary_company_id,
    primary_item_variant_id = EXCLUDED.primary_item_variant_id,
    secondary_item_variant_id = EXCLUDED.secondary_item_variant_id,
    primary_item_unit_id = EXCLUDED.primary_item_unit_id,
    secondary_item_unit_id = EXCLUDED.secondary_item_unit_id,
    updated_at = NOW();

-- name: GetItemVariantMap :one
SELECT a.id,

e.id AS primary_item_id, e.name AS primary_item_name,
b.id AS primary_item_variant_id, b.name AS primary_item_variant_name,
b.price AS primary_item_variant_price,
c.id AS primary_item_unit_id, d.name AS primary_item_unit_name,

i.id AS secondary_item_id, i.name AS secondary_item_name,
f.id AS secondary_item_variant_id, f.name AS secondary_item_variant_name,
f.price AS secondary_item_variant_price,
g.id AS secondary_item_unit_id, h.name AS secondary_item_unit_name

FROM inventory.item_variant_maps a

JOIN inventory.item_variants b ON a.primary_item_variant_id = b.id
JOIN inventory.item_units c ON a.primary_item_unit_id = c.id
JOIN inventory.units d ON c.unit_id = d.id
JOIN inventory.items e ON b.item_id = e.id

JOIN inventory.item_variants f ON a.secondary_item_variant_id = f.id
JOIN inventory.item_units g ON a.secondary_item_unit_id = g.id
JOIN inventory.units h ON g.unit_id = h.id
JOIN inventory.items i ON f.item_id = i.id
WHERE a.id = $1;

-- name: GetItemVariantMaps :many
SELECT a.id,

e.id AS primary_item_id, e.name AS primary_item_name,
b.id AS primary_item_variant_id, b.name AS primary_item_variant_name,
b.price AS primary_item_variant_price,
c.id AS primary_item_unit_id, d.name AS primary_item_unit_name,

i.id AS secondary_item_id, i.name AS secondary_item_name,
f.id AS secondary_item_variant_id, f.name AS secondary_item_variant_name,
f.price AS secondary_item_variant_price,
g.id AS secondary_item_unit_id, h.name AS secondary_item_unit_name

FROM inventory.item_variant_maps a

JOIN inventory.item_variants b ON a.primary_item_variant_id = b.id
JOIN inventory.item_units c ON a.primary_item_unit_id = c.id
JOIN inventory.units d ON c.unit_id = d.id
JOIN inventory.items e ON b.item_id = e.id

JOIN inventory.item_variants f ON a.secondary_item_variant_id = f.id
JOIN inventory.item_units g ON a.secondary_item_unit_id = g.id
JOIN inventory.units h ON g.unit_id = h.id
JOIN inventory.items i ON f.item_id = i.id
WHERE i.company_id = $1 AND e.id = @primary_item_id
ORDER BY b.created_at, c.value;

-- name: UpsertPricelist :one
INSERT INTO inventory.pricelists(id, company_id, name, start_date, end_date, is_default)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO
UPDATE
SET company_id = EXCLUDED.company_id,
    name = EXCLUDED.name,
    start_date = EXCLUDED.start_date,
    end_date = EXCLUDED.end_date,
    is_default = EXCLUDED.is_default,
    updated_at = NOW()
RETURNING *;

-- name: GetPricelists :many
SELECT id, name, start_date, end_date, is_default
FROM inventory.pricelists
WHERE company_id = $1
AND is_deleted = false;

-- name: UpsertPricelistItem :exec
INSERT INTO inventory.pricelist_items(pricelist_id, variant_id, item_unit_id, price)
VALUES ($1, $2, $3, $4) ON CONFLICT (pricelist_id, variant_id, item_unit_id) DO
UPDATE
SET price = EXCLUDED.price,
    updated_at = NOW();

-- name: GetPricelistItems :many
SELECT a.id AS item_id, a.name AS item_name,
b.id AS variant_id, b.name AS variant_name,
c.id AS item_unit_id, d.name AS unit_name,
COALESCE(e.price, 0)::bigint AS price
FROM inventory.items a
JOIN inventory.item_variants b ON a.id = b.item_id
JOIN inventory.item_units c ON b.item_id = c.item_id
JOIN inventory.units d ON c.unit_id = d.id
LEFT JOIN inventory.pricelist_items e ON b.id = e.variant_id
AND c.id = e.item_unit_id AND e.pricelist_id = $2
WHERE a.company_id = $1
AND a.name LIKE $3
AND CASE WHEN b.is_default THEN
NOT EXISTS (
    SELECT
        1
    FROM
        inventory.item_variants a1
    WHERE
        a1.item_id = b.item_id
        AND a1.is_default is false
)
ELSE TRUE END
ORDER BY a.name, b.name, c.value;

-- name: GetPurchaseItems :many
SELECT DISTINCT d.id AS primary_item_id,
d.code AS primary_item_code, d.name AS primary_item_name,
e.id AS secondary_item_id,
e.code AS secondary_item_code, e.name AS secondary_item_name
FROM inventory.item_variant_maps a
JOIN inventory.item_variants b ON a.primary_item_variant_id = b.id
JOIN inventory.item_variants c ON a.secondary_item_variant_id = c.id
JOIN inventory.items d ON b.item_id = d.id
JOIN inventory.items e ON c.item_id = e.id
WHERE a.primary_company_id = $1
AND a.secondary_company_id = $2
AND e.name LIKE $3;

-- name: GetPurchaseItemVariants :many
SELECT DISTINCT a.primary_item_variant_id, b.name AS primary_item_variant_name,
a.secondary_item_variant_id, c.name AS secondary_item_variant_name
FROM inventory.item_variant_maps a
JOIN inventory.item_variants b ON a.primary_item_variant_id = b.id
JOIN inventory.item_variants c ON a.secondary_item_variant_id = c.id
WHERE a.primary_company_id = $1
AND c.item_id = @secondary_item_id
AND c.name LIKE $2;

-- name: GetPurchaseItemVariantUnits :many
SELECT DISTINCT a.primary_item_unit_id, d.name AS primary_item_unit_name,
b.value AS primary_item_unit_value,
a.secondary_item_unit_id, e.name AS secondary_item_unit_name,
c.value AS secondary_item_unit_value, f.price
FROM inventory.item_variant_maps a
JOIN inventory.item_units b ON a.primary_item_unit_id = b.id
JOIN inventory.item_units c ON a.secondary_item_unit_id = c.id
JOIN inventory.units d ON b.unit_id = d.id
JOIN inventory.units e ON c.unit_id = e.id
JOIN inventory.item_variants f ON a.primary_item_variant_id = f.id
WHERE a.primary_company_id = $1
AND a.secondary_item_variant_id = $2
AND e.name LIKE $3;

-- name: GetPOSItems :many
SELECT a.id AS item_id, a.name AS item_name,
b.id AS variant_id, b.name AS variant_name,
c.id AS item_unit_id, d.name AS unit_name,
c.value AS item_unit_value, COALESCE(f.price, 0)::bigint AS price,
SUM(g.amount) AS stock
FROM inventory.items a
JOIN inventory.item_variants b ON a.id = b.item_id
JOIN inventory.item_units c ON b.item_id = c.item_id AND c.value = 1
JOIN inventory.units d ON c.unit_id = d.id
JOIN inventory.pricelists e ON a.company_id = e.company_id
AND e.is_default = true
LEFT JOIN inventory.pricelist_items f ON b.id = f.variant_id
AND c.id = f.item_unit_id AND e.id = f.pricelist_id
JOIN inventory.stock_movements g ON b.id = g.variant_id
AND g.warehouse_id = $2
WHERE a.company_id = $1
AND a.name LIKE $3
AND CASE WHEN b.is_default THEN
NOT EXISTS (
    SELECT
        1
    FROM
        inventory.item_variants a1
    WHERE
        a1.item_id = b.item_id
        AND a1.is_default is false
)
ELSE TRUE END
GROUP BY a.id, b.id, c.id, d.id, f.price
HAVING SUM(g.amount) > 0
ORDER BY a.name, b.name, c.value;

-- name: GetVariantWarehouseRacksByBranch :many
SELECT DISTINCT a.warehouse_id, a.warehouse_rack_id
FROM inventory.stock_movements a
WHERE a.variant_id = $1
AND a.warehouse_id = ANY(@warehouse_ids::text []);

-- name: GetCheckStockHistory :one
SELECT COUNT(a.id)::bigint AS total_count
FROM inventory.update_stocks a
WHERE warehouse_id = ANY(@warehouse_ids::text []);

-- name: GetUnderMinimumOrder :many
SELECT 
    rp.warehouse_id,
	c.id AS item_id,
	c.code AS item_code,
	c.name AS item_name,
	bb.id AS variant_id,
	bb.name AS variant_name,
    e.id AS unit_id,
	e.name AS unit_name,
	a.minimum_stock,
	SUM(rp.amount) AS amount
FROM
	inventory.stock_movements rp
	JOIN inventory.item_variants bb ON bb.id = rp.variant_id
	JOIN inventory.items c ON bb.item_id = c.id
    JOIN inventory.item_units d ON d.item_id = c.id AND d.value = 1
	JOIN inventory.units e ON e.id = d.unit_id
	JOIN inventory.item_reorders a ON a.warehouse_id = rp.warehouse_id
    AND a.variant_id = bb.id 
WHERE rp.company_id = $1 AND rp.branch_id = $2
GROUP BY rp.warehouse_id, bb.id, c.id, e.id, a.minimum_stock
HAVING SUM(rp.amount) < a.minimum_stock
ORDER BY c.name, bb.name;

-- name: GetOutgoingStock :many
SELECT 
	rp.transaction_code AS transaction_code,
	c.id AS item_id,
	c.code AS item_code,
	c.name AS item_name,
	bb.id AS variant_id,
	bb.name AS variant_name,
    e.id AS unit_id,
	e.name AS unit_name,
	-(rp.amount) AS amount
FROM
	inventory.stock_movements rp
	JOIN inventory.item_variants bb ON bb.id = rp.variant_id
	JOIN inventory.items c ON bb.item_id = c.id
    JOIN inventory.item_units d ON d.item_id = c.id AND d.value = 1
	JOIN inventory.units e ON e.id = d.unit_id
WHERE rp.transaction_date BETWEEN @start_date::date AND @end_date::date
AND rp.amount < 0 AND rp.company_id = $1 AND rp.branch_id = $2
ORDER BY rp.transaction_date DESC;

-- name: GetIncomingStock :many
SELECT 
	rp.transaction_code AS transaction_code,
	c.id AS item_id,
	c.code AS item_code,
	c.name AS item_name,
	bb.id AS variant_id,
	bb.name AS variant_name,
    e.id AS unit_id,
	e.name AS unit_name,
	(rp.amount) AS amount
FROM
	inventory.stock_movements rp
	JOIN inventory.item_variants bb ON bb.id = rp.variant_id
	JOIN inventory.items c ON bb.item_id = c.id
    JOIN inventory.item_units d ON d.item_id = c.id AND d.value = 1
	JOIN inventory.units e ON e.id = d.unit_id
WHERE rp.transaction_date BETWEEN @start_date::date AND @end_date::date
AND rp.amount > 0 AND rp.company_id = $1 AND rp.branch_id = $2
ORDER BY rp.transaction_date DESC;

-- name: GetItemHistory :many
SELECT a.transaction_code,
    a.transaction_date,
    item.id AS item_id,
    item.code AS item_code,
    item.name AS item_name,
    variant.id AS variant_id,
    variant.name AS variant_name,
    units.id AS unit_id,
    units.name AS unit_name,
    a.amount
FROM inventory.stock_movements a
    JOIN inventory.item_variants variant ON variant.id = a.variant_id
    JOIN inventory.items item ON item.id = variant.item_id
    JOIN inventory.item_units item_units ON item_units.item_id = item.id AND item_units.value = 1
	JOIN inventory.units units ON units.id = item_units.unit_id
WHERE a.branch_id = $1
    AND a.transaction_date BETWEEN @start_date::date AND @end_date::date
    AND item.id = $2
    AND a.variant_id LIKE $3
ORDER BY a.transaction_date DESC;

-- name: InsertOpeningStock :exec
INSERT INTO inventory.opening_stocks(id,
form_number, transaction_date, company_id, branch_id, warehouse_id, warehouse_rack_id,
variant_id, item_unit_id, item_unit_value, amount, price,
batch, expired_date, item_barcode_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);

-- name: UpsertOpeningStock :exec
INSERT INTO inventory.opening_stocks(id,
form_number, transaction_date, company_id, branch_id, warehouse_id, warehouse_rack_id,
variant_id, item_unit_id, item_unit_value, amount, price,
batch, expired_date, item_barcode_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
ON CONFLICT (id) DO UPDATE SET
form_number = EXCLUDED.form_number,
transaction_date = EXCLUDED.transaction_date,
company_id = EXCLUDED.company_id,
branch_id = EXCLUDED.branch_id,
warehouse_id = EXCLUDED.warehouse_id,
warehouse_rack_id = EXCLUDED.warehouse_rack_id,
variant_id = EXCLUDED.variant_id,
item_unit_id = EXCLUDED.item_unit_id,
item_unit_value = EXCLUDED.item_unit_value,
amount = EXCLUDED.amount,
price = EXCLUDED.price,
batch = EXCLUDED.batch,
expired_date = EXCLUDED.expired_date,
item_barcode_id = EXCLUDED.item_barcode_id;

-- name: DeleteOpeningStock :exec
UPDATE inventory.opening_stocks 
SET 
    is_deleted = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- name: GetOpeningStock :one
SELECT a.id, a.form_number, a.transaction_date, a.warehouse_id, a.warehouse_rack_id,
b.item_id, c.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, e.name AS item_unit_name, a.item_unit_value,
a.amount, a.price, a.batch, a.expired_date
FROM inventory.opening_stocks a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.item_units d ON a.item_unit_id = d.id
JOIN inventory.units e ON d.unit_id = e.id
WHERE a.id = $1;

-- name: GetOpeningStocks :many
SELECT a.id, a.form_number, a.transaction_date, a.warehouse_id, a.warehouse_rack_id,
b.item_id, c.name AS item_name, a.variant_id, b.name AS variant_name,
a.item_unit_id, e.name AS item_unit_name, a.item_unit_value,
a.amount, a.price, a.batch, a.expired_date
FROM inventory.opening_stocks a
JOIN inventory.item_variants b ON a.variant_id = b.id
JOIN inventory.items c ON b.item_id = c.id
JOIN inventory.item_units d ON a.item_unit_id = d.id
JOIN inventory.units e ON d.unit_id = e.id
WHERE a.is_deleted = false
AND a.transaction_date BETWEEN @start_date::date AND @end_date::date
AND a.warehouse_id = ANY(@warehouse_ids::text[]);