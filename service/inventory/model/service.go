package model

import (
	"net/http"
)

type InventoryService interface {
	AddBrand(w http.ResponseWriter, r *http.Request) error
	UpdateBrand(w http.ResponseWriter, r *http.Request) error
	DeleteBrand(w http.ResponseWriter, r *http.Request) error
	GetBrands(w http.ResponseWriter, r *http.Request) error

	AddGroup(w http.ResponseWriter, r *http.Request) error
	UpdateGroup(w http.ResponseWriter, r *http.Request) error
	DeleteGroup(w http.ResponseWriter, r *http.Request) error
	GetGroups(w http.ResponseWriter, r *http.Request) error

	AddUnit(w http.ResponseWriter, r *http.Request) error
	UpdateUnit(w http.ResponseWriter, r *http.Request) error
	GetUnits(w http.ResponseWriter, r *http.Request) error

	AddItem(w http.ResponseWriter, r *http.Request) error
	UpdateItem(w http.ResponseWriter, r *http.Request) error
	GetItems(w http.ResponseWriter, r *http.Request) error

	UpsertItemInfo(w http.ResponseWriter, r *http.Request) error
	GetItemInfo(w http.ResponseWriter, r *http.Request) error

	UpsertItemVariant(w http.ResponseWriter, r *http.Request) error
	GetItemVariants(w http.ResponseWriter, r *http.Request) error
	DeleteItemVariant(w http.ResponseWriter, r *http.Request) error

	UpsertItemUnit(w http.ResponseWriter, r *http.Request) error
	DeleteItemUnit(w http.ResponseWriter, r *http.Request) error
	GetItemUnits(w http.ResponseWriter, r *http.Request) error

	AddInternalStockTransfer(w http.ResponseWriter, r *http.Request) error
	GetInternalStockTransfers(w http.ResponseWriter, r *http.Request) error

	AddUpdateStock(w http.ResponseWriter, r *http.Request) error
	GetUpdateStocks(w http.ResponseWriter, r *http.Request) error

	UpsertItemReorder(w http.ResponseWriter, r *http.Request) error
	GetItemReorders(w http.ResponseWriter, r *http.Request) error

	UpsertUnitCategory(w http.ResponseWriter, r *http.Request) error
	GetUnitCategories(w http.ResponseWriter, r *http.Request) error

	GetVariantWarehouseRacks(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseRackBatches(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseRackBatchExpiredDates(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseRackStock(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseStocks(w http.ResponseWriter, r *http.Request) error

	GetTransferHistory(w http.ResponseWriter, r *http.Request) error
	GetStockHistory(w http.ResponseWriter, r *http.Request) error

	GetItemReorderNotifications(w http.ResponseWriter, r *http.Request) error

	GetSupplierCatalogs(w http.ResponseWriter, r *http.Request) error
	GetMappingItems(w http.ResponseWriter, r *http.Request) error
	GetMappingItemVariants(w http.ResponseWriter, r *http.Request) error
	GetMappingItemUnits(w http.ResponseWriter, r *http.Request) error

	UpsertItemVariantMap(w http.ResponseWriter, r *http.Request) error
	GetItemVariantMaps(w http.ResponseWriter, r *http.Request) error

	UpsertPricelist(w http.ResponseWriter, r *http.Request) error
	GetPricelists(w http.ResponseWriter, r *http.Request) error

	UpsertPricelistItems(w http.ResponseWriter, r *http.Request) error
	GetPricelistItems(w http.ResponseWriter, r *http.Request) error

	GetPurchaseItems(w http.ResponseWriter, r *http.Request) error
	GetPurchaseItemVariants(w http.ResponseWriter, r *http.Request) error
	GetPurchaseItemVariantUnits(w http.ResponseWriter, r *http.Request) error

	GetPOSItems(w http.ResponseWriter, r *http.Request) error
	GetVariantWarehouseRacksByBranch(w http.ResponseWriter, r *http.Request) error

	GetCheckStockHistory(w http.ResponseWriter, r *http.Request) error

	InsertStockMovement(w http.ResponseWriter, r *http.Request) error
	DeleteStockMovement(w http.ResponseWriter, r *http.Request) error

	GetUnderMinimumOrder(w http.ResponseWriter, r *http.Request) error
	GetOutgoingStock(w http.ResponseWriter, r *http.Request) error
	GetIncomingStock(w http.ResponseWriter, r *http.Request) error

	GetItemHistory(w http.ResponseWriter, r *http.Request) error

	AddOpeningStock(w http.ResponseWriter, r *http.Request) error
	GetOpeningStocks(w http.ResponseWriter, r *http.Request) error
	UpsertOpeningStock(w http.ResponseWriter, r *http.Request) error
	DeleteOpeningStock(w http.ResponseWriter, r *http.Request) error
}

type Brand struct {
	CompanyId string `json:"company_id" validate:"required"`
	BrandId   string `json:"brand_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddBrandRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddBrandResponse struct {
	Brand
}

type UpdateBrandRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateBrandResponse struct {
	Brand
}

type DeleteBrandRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeleteBrandResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetBrandsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Group struct {
	CompanyId string `json:"company_id" validate:"required"`
	GroupId   string `json:"group_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddGroupRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type AddGroupResponse struct {
	Group
}

type UpdateGroupRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateGroupResponse struct {
	Group
}

type DeleteGroupRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeleteGroupResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetGroupsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type Unit struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	UnitId         string `json:"unit_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type AddUnitRequest struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type AddUnitResponse struct {
	Unit
}

type UpdateUnitRequest struct {
	Id             string `json:"id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
}

type UpdateUnitResponse struct {
	Unit
}

type GetUnitsRequest struct {
	CompanyId      string `json:"company_id" validate:"required"`
	UnitCategoryId string `json:"unit_category_id"`
	Keyword        string `json:"keyword"`
}

type ItemGroup struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type Item struct {
	CompanyId   string      `json:"company_id" validate:"required"`
	ItemId      string      `json:"item_id" validate:"required"`
	VariantId   string      `json:"variant_id" validate:"required"`
	ImageUrl    string      `json:"image_url" validate:"required"`
	Code        string      `json:"code" validate:"required"`
	Barcode     string      `json:"barcode" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	VariantName string      `json:"variant_name" validate:"required"`
	BrandId     string      `json:"brand_id" validate:"required"`
	BrandName   string      `json:"brand_name" validate:"required"`
	Groups      []ItemGroup `json:"groups" validate:"required"`
	Tag         []string    `json:"tag" validate:"required"`
	Description string      `json:"description" validate:"required"`
	IsDefault   bool        `json:"is_default" validate:"required"`
	Price       string      `json:"price" validate:"required"`
}

type ItemWithStock struct {
	CompanyId     string      `json:"company_id" validate:"required"`
	ItemId        string      `json:"item_id" validate:"required"`
	VariantId     string      `json:"variant_id" validate:"required"`
	ImageUrl      string      `json:"image_url" validate:"required"`
	Code          string      `json:"code" validate:"required"`
	Barcode       string      `json:"barcode" validate:"required"`
	Name          string      `json:"name" validate:"required"`
	VariantName   string      `json:"variant_name" validate:"required"`
	BrandId       string      `json:"brand_id" validate:"required"`
	BrandName     string      `json:"brand_name" validate:"required"`
	Groups        []ItemGroup `json:"groups" validate:"required"`
	Tag           []string    `json:"tag" validate:"required"`
	Description   string      `json:"description" validate:"required"`
	IsDefault     bool        `json:"is_default" validate:"required"`
	Price         string      `json:"price" validate:"required"`
	AmountInStock *string     `json:"amount_in_stock" validate:"required"`
}

type AddItemRequest struct {
	CompanyId   string   `json:"company_id" validate:"required"`
	ImageUrl    string   `json:"image_url"`
	Code        string   `json:"code"`
	Name        string   `json:"name" validate:"required"`
	Barcode     string   `json:"barcode"`
	BrandId     string   `json:"brand_id"`
	GroupIds    []string `json:"group_ids" validate:"required"`
	Tag         []string `json:"tag"`
	Price       string   `json:"price" validate:"required"`
	Description string   `json:"description"`
}

type AddItemResponse struct {
	Item
}

type UpdateItemRequest struct {
	ItemId      string   `json:"item_id" validate:"required"`
	ImageUrl    string   `json:"image_url"`
	Barcode     string   `json:"barcode"`
	Name        string   `json:"name" validate:"required"`
	BrandId     string   `json:"brand_id"`
	GroupIds    []string `json:"group_ids" validate:"required"`
	Tag         []string `json:"tag"`
	Price       string   `json:"price" validate:"required"`
	Description string   `json:"description"`
}

type UpdateItemResponse struct {
	Item
}

type GetItemsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetItemsResponse struct {
	Items []ItemWithStock `json:"items" validate:"required"`
}

type ItemInfo struct {
	ItemId                     string `json:"item_id" validate:"required"`
	IsPurchase                 bool   `json:"is_purchase" validate:"required"`
	IsSale                     bool   `json:"is_sale" validate:"required"`
	IsRawMaterial              bool   `json:"is_raw_material" validate:"required"`
	IsAsset                    bool   `json:"is_asset" validate:"required"`
	PurchaseChartOfAccountId   string `json:"purchase_chart_of_account_id" validate:"required"`
	PurchaseChartOfAccountName string `json:"purchase_chart_of_account_name" validate:"required"`
	SaleChartOfAccountId       string `json:"sale_chart_of_account_id" validate:"required"`
	SaleChartOfAccountName     string `json:"sale_chart_of_account_name" validate:"required"`
	PurchaseItemUnitId         string `json:"purchase_item_unit_id" validate:"required"`
	PurchaseItemUnitName       string `json:"purchase_item_unit_name" validate:"required"`
}

type UpsertItemInfoRequest struct {
	ItemId                   string `json:"item_id" validate:"required"`
	IsPurchase               bool   `json:"is_purchase"`
	IsSale                   bool   `json:"is_sale"`
	IsRawMaterial            bool   `json:"is_raw_material"`
	IsAsset                  bool   `json:"is_asset"`
	PurchaseChartOfAccountId string `json:"purchase_chart_of_account_id" validate:"required"`
	SaleChartOfAccountId     string `json:"sale_chart_of_account_id" validate:"required"`
	PurchaseItemUnitId       string `json:"purchase_item_unit_id" validate:"required"`
}

type UpsertItemInfoResponse struct {
	ItemInfo
}

type GetItemInfoRequest struct {
	ItemId string `json:"item_id" validate:"required"`
}

type GetItemInfoResponse struct {
	ItemInfo
}

type UpsertItemVariantRequest struct {
	ItemVariantId string `json:"item_variant_id"`
	ItemId        string `json:"item_id" validate:"required"`
	ImageUrl      string `json:"image_url"`
	Barcode       string `json:"barcode"`
	Name          string `json:"name" validate:"required"`
	Price         string `json:"price" validate:"required"`
}

type UpsertItemVariantResponse struct {
	Item
}

type GetItemVariantsRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id"`
	Keyword string `json:"keyword"`
}

type GetItemVariantsResponse struct {
	ItemVariants []Item `json:"item_variants" validate:"required"`
}

type DeleteItemVariantRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeleteItemVariantResponse struct {
	Message string `json:"message" validate:"required"`
}

type ItemUnit struct {
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	ItemId     string `json:"item_id" validate:"required"`
	UnitId     string `json:"unit_id" validate:"required"`
	UnitName   string `json:"unit_name" validate:"required"`
	Value      string `json:"value" validate:"required"`
	IsDefault  bool   `json:"is_default" validate:"required"`
}

type UpsertItemUnitRequest struct {
	ItemUnitId string `json:"item_unit_id"`
	ItemId     string `json:"item_id" validate:"required"`
	UnitId     string `json:"unit_id" validate:"required"`
	Value      string `json:"value" validate:"required"`
	IsDefault  bool   `json:"is_default"`
}

type UpsertItemUnitResponse struct {
	ItemUnit
}

type DeleteItemUnitRequest struct {
	ItemUnitId string `json:"item_unit_id"`
}

type DeleteItemUnitResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetItemUnitsRequest struct {
	Id      string `json:"id"`
	ItemId  string `json:"item_id" validate:"required"`
	Keyword string `json:"keyword"`
}

type GetItemUnitsResponse struct {
	ItemUnits []ItemUnit `json:"item_units" validate:"required"`
}

type InternalStockTransferItem struct {
	DetailId          string  `json:"detail_id" validate:"required"`
	WarehouseRackId   string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string  `json:"warehouse_rack_name" validate:"required"`
	ItemName          string  `json:"item_name" validate:"required"`
	VariantId         string  `json:"variant_id" validate:"required"`
	VariantName       string  `json:"variant_name" validate:"required"`
	ItemUnitId        string  `json:"item_unit_id" validate:"required"`
	ItemUnitName      string  `json:"item_unit_name" validate:"required"`
	ItemUnitValue     string  `json:"item_unit_value" validate:"required"`
	Amount            string  `json:"amount" validate:"required"`
	Batch             *string `json:"batch" validate:"required"`
	ExpiredDate       *string `json:"expired_date" validate:"required"`
}

type InternalStockTransfer struct {
	TransactionId            string                      `json:"transaction_id" validate:"required"`
	SourceWarehouseId        string                      `json:"source_warehouse_id" validate:"required"`
	SourceWarehouseName      string                      `json:"source_warehouse_name" validate:"required"`
	DestinationWarehouseId   string                      `json:"destination_warehouse_id" validate:"required"`
	DestinationWarehouseName string                      `json:"destination_warehouse_name" validate:"required"`
	FormNumber               string                      `json:"form_number" validate:"required"`
	TransactionDate          string                      `json:"transaction_date" validate:"required"`
	Items                    []InternalStockTransferItem `json:"items" validate:"required"`
}

type InternalStockTransferItemRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id"`
	VariantId       string `json:"variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type AddInternalStockTransferRequest struct {
	CompanyId              string                             `json:"company_id" validate:"required"`
	BranchId               string                             `json:"branch_id" validate:"required"`
	SourceWarehouseId      string                             `json:"source_warehouse_id" validate:"required"`
	DestinationWarehouseId string                             `json:"destination_warehouse_id" validate:"required"`
	TransactionDate        string                             `json:"transaction_date" validate:"required"`
	Items                  []InternalStockTransferItemRequest `json:"items" validate:"required"`
}

type AddInternalStockTransferResponse struct {
	InternalStockTransfer
}

type GetInternalStockTransfersRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetInternalStockTransfersResponse struct {
	InternalStockTransfers []InternalStockTransfer `json:"internal_stock_transfers" validate:"required"`
}

type UpdateStock struct {
	TransactionId     string  `json:"transaction_id" validate:"required"`
	FormNumber        string  `json:"form_number" validate:"required"`
	TransactionDate   string  `json:"transaction_date" validate:"required"`
	WarehouseId       string  `json:"warehouse_id" validate:"required"`
	WarehouseName     string  `json:"warehouse_name" validate:"required"`
	WarehouseRackId   string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string  `json:"warehouse_rack_name" validate:"required"`
	ItemId            string  `json:"item_id" validate:"required"`
	ItemName          string  `json:"item_name" validate:"required"`
	VariantId         string  `json:"variant_id" validate:"required"`
	VariantName       string  `json:"variant_name" validate:"required"`
	ItemUnitId        string  `json:"item_unit_id" validate:"required"`
	ItemUnitName      string  `json:"item_unit_name" validate:"required"`
	ItemUnitValue     string  `json:"item_unit_value" validate:"required"`
	BeginningStock    string  `json:"beginning_stock" validate:"required"`
	EndingStock       string  `json:"ending_stock" validate:"required"`
	Batch             *string `json:"batch" validate:"required"`
	ExpiredDate       *string `json:"expired_date" validate:"required"`
}

type AddUpdateStockRequest struct {
	TransactionDate string `json:"transaction_date" validate:"required"`
	CompanyId       string `json:"company_id" validate:"required"`
	BranchId        string `json:"branch_id" validate:"required"`
	WarehouseId     string `json:"warehouse_id" validate:"required"`
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	BeginningStock  string `json:"beginning_stock" validate:"required"`
	EndingStock     string `json:"ending_stock" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type AddUpdateStockResponse struct {
	UpdateStock
}

type GetUpdateStocksRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetUpdateStocksResponse struct {
	UpdateStocks []UpdateStock `json:"update_stocks" validate:"required"`
}

type ItemReorder struct {
	Id            string `json:"id" validate:"required"`
	ItemId        string `json:"item_id" validate:"required"`
	ItemName      string `json:"item_name" validate:"required"`
	VariantId     string `json:"variant_id" validate:"required"`
	VariantName   string `json:"variant_name" validate:"required"`
	ItemUnitId    string `json:"item_unit_id" validate:"required"`
	ItemUnitName  string `json:"item_unit_name" validate:"required"`
	WarehouseId   string `json:"warehouse_id" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
	MinimumStock  string `json:"minimum_stock" validate:"required"`
}

type UpsertItemReorderRequest struct {
	Id           string `json:"id"`
	VariantId    string `json:"variant_id" validate:"required"`
	ItemUnitId   string `json:"item_unit_id" validate:"required"`
	WarehouseId  string `json:"warehouse_id" validate:"required"`
	MinimumStock string `json:"minimum_stock" validate:"required"`
}

type UpsertItemReorderResponse struct {
	ItemReorder ItemReorder `json:"item_reorder" validate:"required"`
}

type GetItemReordersRequest struct {
	ItemId      string `json:"item_id"`
	WarehouseId string `json:"warehouse_id"`
}

type GetItemReordersResponse struct {
	ItemReorders []ItemReorder `json:"item_reorders" validate:"required"`
}

type UnitCategory struct {
	Id        string `json:"id" validate:"required"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type UpsertUnitCategoryRequest struct {
	Id        string `json:"id"`
	CompanyId string `json:"company_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

type UpsertUnitCategoryResponse struct {
	UnitCategory
}

type GetUnitCategoriesRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	Keyword   string `json:"keyword"`
}

type GetUnitCategoriesResponse struct {
	UnitCategories []UnitCategory `json:"unit_categories" validate:"required"`
}

type WarehouseRack struct {
	RackId      string `json:"rack_id" validate:"required"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type GetVariantWarehouseRacksRequest struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
}

type GetVariantWarehouseRacksResponse struct {
	WarehouseRacks []WarehouseRack `json:"warehouse_racks" validate:"required"`
}

type GetVariantWarehouseRackBatchesRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
}

type GetVariantWarehouseRackBatchesResponse struct {
	Batches []*string `json:"batches" validate:"required"`
}

type GetVariantWarehouseRackBatchExpiredDatesRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	Batch           string `json:"batch"`
}

type GetVariantWarehouseRackBatchExpiredDatesResponse struct {
	ExpiredDates []*string `json:"expired_dates" validate:"required"`
}

type GetVariantWarehouseRackStockRequest struct {
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type GetVariantWarehouseRackStockResponse struct {
	ItemBarcodeId string `json:"item_barcode_id" validate:"required"`
	Stock         string `json:"stock" validate:"required"`
}

type VariantStock struct {
	ItemId      string `json:"item_id" validate:"required"`
	ItemName    string `json:"item_name" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
	VariantName string `json:"variant_name" validate:"required"`
	Stock       string `json:"stock" validate:"required"`
}

type GetVariantWarehouseStocksRequest struct {
	WarehouseId string `json:"warehouse_id" validate:"required"`
}

type GetVariantWarehouseStocksResponse struct {
	VariantStocks []VariantStock `json:"variant_stocks" validate:"required"`
}

type TransferHistory struct {
	FormNumber               string `json:"form_number" validate:"required"`
	TransactionDate          string `json:"transaction_date" validate:"required"`
	SourceWarehouseId        string `json:"source_warehouse_id" validate:"required"`
	SourceWarehouseName      string `json:"source_warehouse_name" validate:"required"`
	DestinationWarehouseId   string `json:"destination_warehouse_id" validate:"required"`
	DestinationWarehouseName string `json:"destination_warehouse_name" validate:"required"`
	ItemId                   string `json:"item_id" validate:"required"`
	ItemName                 string `json:"item_name" validate:"required"`
	ItemImageUrl             string `json:"item_image_url" validate:"required"`
	VariantId                string `json:"variant_id" validate:"required"`
	VariantName              string `json:"variant_name" validate:"required"`
	Amount                   string `json:"amount" validate:"required"`
}

type GetTransferHistoryRequest struct {
	BranchId         string `json:"branch_id" validate:"required"`
	WarehouseId      string `json:"warehouse_id"`
	ItemId           string `json:"item_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	IsReceivedFilter *bool  `json:"is_received_filter"`
}

type GetTransferHistoryResponse struct {
	TransferHistories []TransferHistory `json:"transfer_histories" validate:"required"`
}

type StockHistory struct {
	FormNumber      string `json:"form_number" validate:"required"`
	TransactionDate string `json:"transaction_date" validate:"required"`
	ItemId          string `json:"item_id" validate:"required"`
	ItemName        string `json:"item_name" validate:"required"`
	ItemImageUrl    string `json:"item_image_url" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	VariantName     string `json:"variant_name" validate:"required"`
	Onhand          string `json:"onhand" validate:"required"`
	Calculated      string `json:"calculated" validate:"required"`
}

type GetStockHistoryRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	ItemId    string `json:"item_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetStockHistoryResponse struct {
	StockHistories []StockHistory `json:"stock_histories" validate:"required"`
}

type ItemReorderNotification struct {
	ItemId       string  `json:"item_id" validate:"required"`
	ItemName     string  `json:"item_name" validate:"required"`
	VariantId    string  `json:"variant_id" validate:"required"`
	VariantName  string  `json:"variant_name" validate:"required"`
	CurrentStock string  `json:"current_stock" validate:"required"`
	MinimumStock *string `json:"minimum_stock" validate:"required"`
}

type GetItemReorderNotificationsRequest struct {
	WarehouseId string   `json:"warehouse_id" validate:"required"`
	VariantIds  []string `json:"variant_ids" validate:"required"`
}

type GetItemReorderNotificationsResponse struct {
	ItemVariants []ItemReorderNotification `json:"item_variants" validate:"required"`
}

type SupplierCatalog struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemId      string `json:"item_id" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Barcode     string `json:"barcode" validate:"required"`
	Name        string `json:"name" validate:"required"`
	VariantName string `json:"variant_name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	BrandName   string `json:"brand_name" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
	Price       string `json:"price" validate:"required"`
	IsMapped    bool   `json:"is_mapped" validate:"required"`
}

type GetSupplierCatalogsRequest struct {
	PrimaryCompanyId   string `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	Keyword            string `json:"keyword"`
}

type GetSupplierCatalogsResponse struct {
	SupplierCatalogs []SupplierCatalog `json:"supplier_catalogs" validate:"required"`
}

type MappingItem struct {
	ItemId string `json:"item_id" validate:"required"`
	Code   string `json:"code" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type GetMappingItemsRequest struct {
	PrimaryCompanyId   string `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	ItemId             string `json:"item_id" validate:"required"`
	Keyword            string `json:"keyword"`
}

type GetMappingItemsResponse struct {
	Items []MappingItem `json:"items" validate:"required"`
}

type MappingItemVariant struct {
	CompanyId   string `json:"company_id" validate:"required"`
	ItemId      string `json:"item_id" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
	ImageUrl    string `json:"image_url" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Barcode     string `json:"barcode" validate:"required"`
	Name        string `json:"name" validate:"required"`
	VariantName string `json:"variant_name" validate:"required"`
	BrandId     string `json:"brand_id" validate:"required"`
	BrandName   string `json:"brand_name" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
	Price       string `json:"price" validate:"required"`
}

type GetMappingItemVariantsRequest struct {
	PrimaryCompanyId   string `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	ItemId             string `json:"item_id" validate:"required"`
	Keyword            string `json:"keyword"`
}

type GetMappingItemVariantsResponse struct {
	ItemVariants []MappingItemVariant `json:"item_variants" validate:"required"`
}

type MappingItemUnit struct {
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	UnitName   string `json:"unit_name" validate:"required"`
}

type GetMappingItemUnitsRequest struct {
	PrimaryCompanyId   string `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	VariantId          string `json:"variant_id" validate:"required"`
	IsSupplier         bool   `json:"is_supplier"`
}

type GetMappingItemUnitsResponse struct {
	ItemUnits []MappingItemUnit `json:"item_units" validate:"required"`
}

type ItemVariantMap struct {
	ItemVariantMapId          string `json:"item_variant_map_id" validate:"required"`
	PrimaryItemId             string `json:"primary_item_id" validate:"required"`
	PrimaryItemName           string `json:"primary_item_name" validate:"required"`
	PrimaryItemVariantId      string `json:"primary_item_variant_id" validate:"required"`
	PrimaryItemVariantName    string `json:"primary_item_variant_name" validate:"required"`
	PrimaryItemVariantPrice   string `json:"primary_item_variant_price" validate:"required"`
	PrimaryItemUnitId         string `json:"primary_item_unit_id" validate:"required"`
	PrimaryItemUnitName       string `json:"primary_item_unit_name" validate:"required"`
	SecondaryItemId           string `json:"secondary_item_id" validate:"required"`
	SecondaryItemName         string `json:"secondary_item_name" validate:"required"`
	SecondaryItemVariantId    string `json:"secondary_item_variant_id" validate:"required"`
	SecondaryItemVariantName  string `json:"secondary_item_variant_name" validate:"required"`
	SecondaryItemVariantPrice string `json:"secondary_item_variant_price" validate:"required"`
	SecondaryItemUnitId       string `json:"secondary_item_unit_id" validate:"required"`
	SecondaryItemUnitName     string `json:"secondary_item_unit_name" validate:"required"`
}

type UpsertItemVariantMapRequest struct {
	ItemVariantMapId       string `json:"item_variant_map_id"`
	PrimaryCompanyId       string `json:"primary_company_id" validate:"required"`
	PrimaryItemVariantId   string `json:"primary_item_variant_id" validate:"required"`
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	SecondaryCompanyId     string `json:"secondary_company_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id" validate:"required"`
}

type UpsertItemVariantMapResponse struct {
	ItemVariantMap
}

type GetItemVariantMapsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	ItemId    string `json:"item_id" validate:"required"`
}

type GetItemVariantMapsResponse struct {
	ItemVariantMaps []ItemVariantMap `json:"item_variant_mappings" validate:"required"`
}

type Pricelist struct {
	Id        string  `json:"pricelist_id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	StartDate string  `json:"start_date" validate:"required"`
	EndDate   *string `json:"end_date" validate:"required"`
	IsDefault bool    `json:"is_default" validate:"required"`
}

type UpsertPricelistRequest struct {
	PricelistId string `json:"pricelist_id"`
	CompanyId   string `json:"company_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date"`
	IsDefault   bool   `json:"is_default"`
}

type UpsertPricelistResponse struct {
	Pricelist
}

type GetPricelistsRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
}

type GetPricelistsResponse struct {
	Pricelists []Pricelist `json:"pricelists" validate:"required"`
}

type PricelistItem struct {
	ItemId      string `json:"item_id" validate:"required"`
	ItemName    string `json:"item_name" validate:"required"`
	VariantId   string `json:"variant_id" validate:"required"`
	VariantName string `json:"variant_name" validate:"required"`
	ItemUnitId  string `json:"item_unit_id" validate:"required"`
	UnitName    string `json:"unit_name" validate:"required"`
	Price       string `json:"price" validate:"required"`
}

type PricelistItemRequest struct {
	VariantId  string `json:"variant_id" validate:"required"`
	ItemUnitId string `json:"item_unit_id" validate:"required"`
	Price      string `json:"price" validate:"required"`
}

type UpsertPricelistItemsRequest struct {
	PricelistId    string                 `json:"pricelist_id" validate:"required"`
	PricelistItems []PricelistItemRequest `json:"pricelist_items" validate:"required"`
}

type UpsertPricelistItemsResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetPricelistItemsRequest struct {
	CompanyId   string `json:"company_id" validate:"required"`
	PricelistId string `json:"pricelist_id" validate:"required"`
	Keyword     string `json:"keyword"`
}

type GetPricelistItemsResponse struct {
	PricelistItems []PricelistItem `json:"pricelist_items" validate:"required"`
}

type PurchaseItem struct {
	PrimaryItemId     string `json:"primary_item_id" validate:"required"`
	PrimaryItemCode   string `json:"primary_item_code" validate:"required"`
	PrimaryItemName   string `json:"primary_item_name" validate:"required"`
	SecondaryItemId   string `json:"secondary_item_id" validate:"required"`
	SecondaryItemCode string `json:"secondary_item_code" validate:"required"`
	SecondaryItemName string `json:"secondary_item_name" validate:"required"`
}

type GetPurchaseItemsRequest struct {
	PrimaryCompanyId   string `json:"primary_company_id" validate:"required"`
	SecondaryCompanyId string `json:"secondary_company_id" validate:"required"`
	Keyword            string `json:"keyword"`
}

type GetPurchaseItemsResponse struct {
	PurchaseItems []PurchaseItem `json:"purchase_items" validate:"required"`
}

type PurchaseItemVariant struct {
	PrimaryItemVariantId     string `json:"primary_item_variant_id" validate:"required"`
	PrimaryItemVariantName   string `json:"primary_item_variant_name" validate:"required"`
	SecondaryItemVariantId   string `json:"secondary_item_variant_id" validate:"required"`
	SecondaryItemVariantName string `json:"secondary_item_variant_name" validate:"required"`
}

type GetPurchaseItemVariantsRequest struct {
	PrimaryCompanyId string `json:"primary_company_id" validate:"required"`
	SecondaryItemId  string `json:"secondary_item_id" validate:"required"`
	Keyword          string `json:"keyword"`
}

type GetPurchaseItemVariantsResponse struct {
	PurchaseItemVariants []PurchaseItemVariant `json:"purchase_item_variants" validate:"required"`
}

type PurchaseItemVariantUnit struct {
	PrimaryItemUnitId      string `json:"primary_item_unit_id" validate:"required"`
	PrimaryItemUnitName    string `json:"primary_item_unit_name" validate:"required"`
	PrimaryItemUnitValue   string `json:"primary_item_unit_value" validate:"required"`
	SecondaryItemUnitId    string `json:"secondary_item_unit_id" validate:"required"`
	SecondaryItemUnitName  string `json:"secondary_item_unit_name" validate:"required"`
	SecondaryItemUnitValue string `json:"secondary_item_unit_value" validate:"required"`
	Price                  string `json:"price" validate:"required"`
}

type GetPurchaseItemVariantUnitsRequest struct {
	PrimaryCompanyId       string `json:"primary_company_id" validate:"required"`
	SecondaryItemVariantId string `json:"secondary_item_variant_id" validate:"required"`
	Keyword                string `json:"keyword"`
}

type GetPurchaseItemVariantUnitsResponse struct {
	PurchaseItemVariantUnits []PurchaseItemVariantUnit `json:"purchase_item_variant_units" validate:"required"`
}

type POSItem struct {
	ItemId        string `json:"item_id" validate:"required"`
	ItemName      string `json:"item_name" validate:"required"`
	VariantId     string `json:"variant_id" validate:"required"`
	VariantName   string `json:"variant_name" validate:"required"`
	ItemUnitId    string `json:"item_unit_id" validate:"required"`
	UnitName      string `json:"unit_name" validate:"required"`
	ItemUnitValue string `json:"item_unit_value" validate:"required"`
	Price         string `json:"price" validate:"required"`
}

type GetPOSItemsRequest struct {
	CompanyId   string `json:"company_id" validate:"required"`
	WarehouseId string `json:"warehouse_id" validate:"required"`
	Keyword     string `json:"keyword"`
}

type GetPOSItemsResponse struct {
	POSItems []POSItem `json:"pos_items" validate:"required"`
}

type WarehouseAndRack struct {
	WarehouseId   string `json:"warehouse_id" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
	RackId        string `json:"rack_id" validate:"required"`
	RackName      string `json:"rack_name" validate:"required"`
}

type GetVariantWarehouseRacksByBranchRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	VariantId string `json:"variant_id" validate:"required"`
}

type GetVariantWarehouseRacksByBranchResponse struct {
	WarehouseAndRacks []WarehouseAndRack `json:"warehouse_racks" validate:"required"`
}

type GetCheckStockHistoryRequest struct {
	BranchIds []string `json:"branch_ids" validate:"required"`
}

type GetCheckStockHistoryResponse struct {
	Status bool `json:"status" validate:"required"`
}

type InsertStockMovementRequest struct {
	TransactionId        string `json:"transaction_id" validate:"required"`
	CompanyId            string `json:"company_id" validate:"required"`
	BranchId             string `json:"branch_id" validate:"required"`
	TransactionCode      string `json:"transaction_code" validate:"required"`
	TransactionDate      string `json:"transaction_date" validate:"required"`
	TransactionReference string `json:"transaction_reference" validate:"required"`
	DetailTransactionId  string `json:"detail_transaction_id" validate:"required"`
	WarehouseId          string `json:"warehouse_id" validate:"required"`
	WarehouseRackId      string `json:"warehouse_rack_id" validate:"required"`
	VariantId            string `json:"variant_id" validate:"required"`
	ItemBarcodeId        string `json:"item_barcode_id" validate:"required"`
	Amount               string `json:"amount" validate:"required"`
}

type InsertStockMovementResponse struct {
	Message string `json:"message" validate:"required"`
}

type DeleteStockMovementRequest struct {
	TransactionId        string `json:"transaction_id" validate:"required"`
	TransactionReference string `json:"transaction_reference" validate:"required"`
}

type DeleteStockMovementResponse struct {
	Message string `json:"message" validate:"required"`
}

type GetUnderMinimumOrderRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
}

type GetUnderMinimumOrderResponseStruct struct {
	WarehouseName string `json:"warehouse_name" validate:"required"`
	ItemId        string `json:"item_id" validate:"required"`
	ItemCode      string `json:"item_code" validate:"required"`
	ItemName      string `json:"item_name" validate:"required"`
	VariantId     string `json:"variant_id" validate:"required"`
	VariantName   string `json:"variant_name" validate:"required"`
	UnitId        string `json:"unit_id" validate:"required"`
	UnitName      string `json:"unit_name" validate:"required"`
	MinimumStock  string `json:"minimum_stock" validate:"required"`
	Amount        string `json:"amount" validate:"required"`
}

type GetUnderMinimumOrderResponse struct {
	UnderMinimumOrder []GetUnderMinimumOrderResponseStruct `json:"under_minimum_orders" validate:"required"`
}

type GetOutgoingStockRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type GetOutgoingStockResponseStruct struct {
	TransactionCode string `json:"transaction_code" validate:"required"`
	ItemId          string `json:"item_id" validate:"required"`
	ItemCode        string `json:"item_code" validate:"required"`
	ItemName        string `json:"item_name" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	VariantName     string `json:"variant_name" validate:"required"`
	UnitId          string `json:"unit_id" validate:"required"`
	UnitName        string `json:"unit_name" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
}

type GetOutgoingStockResponse struct {
	OutgoingStock []GetOutgoingStockResponseStruct `json:"outgoing_stocks" validate:"required"`
}

type GetIncomingStockRequest struct {
	CompanyId string `json:"company_id" validate:"required"`
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type GetIncomingStockResponseStruct struct {
	TransactionCode string `json:"transaction_code" validate:"required"`
	ItemId          string `json:"item_id" validate:"required"`
	ItemCode        string `json:"item_code" validate:"required"`
	ItemName        string `json:"item_name" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	VariantName     string `json:"variant_name" validate:"required"`
	UnitId          string `json:"unit_id" validate:"required"`
	UnitName        string `json:"unit_name" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
}

type GetIncomingStockResponse struct {
	IncomingStock []GetIncomingStockResponseStruct `json:"incoming_stocks" validate:"required"`
}

type ItemHistory struct {
	TransactionCode string `json:"transaction_code" validate:"required"`
	TransactionDate string `json:"transaction_date" validate:"required"`
	ItemId          string `json:"item_id" validate:"required"`
	ItemCode        string `json:"item_code" validate:"required"`
	ItemName        string `json:"item_name" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	VariantName     string `json:"variant_name" validate:"required"`
	UnitId          string `json:"unit_id" validate:"required"`
	UnitName        string `json:"unit_name" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
}

type GetItemHistoryRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	ItemId    string `json:"item_id" validate:"required"`
	VariantId string `json:"variant_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetItemHistoryResponse struct {
	ItemHistories []ItemHistory `json:"item_histories" validate:"required"`
}

type OpeningStock struct {
	TransactionId     string  `json:"transaction_id" validate:"required"`
	FormNumber        string  `json:"form_number" validate:"required"`
	TransactionDate   string  `json:"transaction_date" validate:"required"`
	WarehouseId       string  `json:"warehouse_id" validate:"required"`
	WarehouseName     string  `json:"warehouse_name" validate:"required"`
	WarehouseRackId   string  `json:"warehouse_rack_id" validate:"required"`
	WarehouseRackName string  `json:"warehouse_rack_name" validate:"required"`
	ItemId            string  `json:"item_id" validate:"required"`
	ItemName          string  `json:"item_name" validate:"required"`
	VariantId         string  `json:"variant_id" validate:"required"`
	VariantName       string  `json:"variant_name" validate:"required"`
	ItemUnitId        string  `json:"item_unit_id" validate:"required"`
	ItemUnitName      string  `json:"item_unit_name" validate:"required"`
	ItemUnitValue     string  `json:"item_unit_value" validate:"required"`
	Amount            string  `json:"amount" validate:"required"`
	Price             string  `json:"price" validate:"required"`
	Batch             *string `json:"batch" validate:"required"`
	ExpiredDate       *string `json:"expired_date" validate:"required"`
}

type AddOpeningStockRequest struct {
	TransactionDate string `json:"transaction_date" validate:"required"`
	CompanyId       string `json:"company_id" validate:"required"`
	BranchId        string `json:"branch_id" validate:"required"`
	WarehouseId     string `json:"warehouse_id" validate:"required"`
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
	Price           string `json:"price" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type AddOpeningStockResponse struct {
	OpeningStock
}

type GetOpeningStocksRequest struct {
	BranchId  string `json:"branch_id" validate:"required"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetOpeningStocksResponse struct {
	OpeningStocks []OpeningStock `json:"opening_stocks" validate:"required"`
}

type UpsertOpeningStockRequest struct {
	Id              string `json:"id"`
	TransactionDate string `json:"transaction_date" validate:"required"`
	CompanyId       string `json:"company_id" validate:"required"`
	BranchId        string `json:"branch_id" validate:"required"`
	WarehouseId     string `json:"warehouse_id" validate:"required"`
	WarehouseRackId string `json:"warehouse_rack_id" validate:"required"`
	VariantId       string `json:"variant_id" validate:"required"`
	ItemUnitId      string `json:"item_unit_id" validate:"required"`
	ItemUnitValue   string `json:"item_unit_value" validate:"required"`
	Amount          string `json:"amount" validate:"required"`
	Price           string `json:"price" validate:"required"`
	Batch           string `json:"batch"`
	ExpiredDate     string `json:"expired_date"`
}

type UpsertOpeningStockResponse struct {
	OpeningStock
}

type DeleteOpeningStockRequest struct {
	Id string `json:"id" validate:"required"`
}

type DeleteOpeningStockResponse struct {
	Message string `json:"message" validate:"required"`
}
