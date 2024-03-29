package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/purchasing/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	UpsertPurchaseOrderPath      = "/purchasing/order/upsert"
	UpsertPurchaseOrderItemPath  = "/purchasing/order/item/upsert"
	UpdatePurchaseOrderItemsPath = "/purchasing/order/items/update"
	GetPurchaseOrdersPath        = "/purchasing/orders"
	GetPurcaseOrderItemsPath     = "/purchasing/order/items"
	UpdatePurchaseSettingPath    = "/purchasing/setting/update"
	GetPurchaseSettingPath       = "/purchasing/setting"

	GetCheckPurchaseOrdersPath = "/purchasing/orders/check"

	UpdatePurchaseOrderStatusPath = "/purchasing/order/status"

	UpsertReceiptOrderPath       = "/purchasing/receipt-order/upsert"
	GetReceiptOrdersPath         = "/purchasing/receipt-orders"
	UpdateReceiptOrderItemsPath  = "/purchasing/receipt-order/items/update"
	GetReceiptOrderItemsPath     = "/purchasing/receipt-order/items"
	DeleteReceiptOrderPath       = "/purchasing/receipt-order/delete"
	UpdateReceiptOrderStatusPath = "/purchasing/receipt-order/status"

	UpsertPurchaseInvoicePath   = "/purchasing/invoice/upsert"
	GetPurchaseInvoicesPath     = "/purchasing/invoices"
	GetPurchaseInvoiceItemsPath = "/purchasing/invoice/items"
)

func (c *component) Routes(purchasingService model.PurchasingService) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Method("POST", UpsertPurchaseOrderPath, httpHandler.New(purchasingService.UpsertPurchaseOrder))
	mux.Method("POST", UpsertPurchaseOrderItemPath, httpHandler.New(purchasingService.UpsertPurchaseOrderItem))
	mux.Method("POST", UpdatePurchaseOrderItemsPath, httpHandler.New(purchasingService.UpdatePurchaseOrderItems))
	mux.Method("POST", GetPurchaseOrdersPath, httpHandler.New(purchasingService.GetPurchaseOrders))
	mux.Method("POST", GetPurcaseOrderItemsPath, httpHandler.New(purchasingService.GetPurchaseOrderItems))

	mux.Method("POST", UpdatePurchaseSettingPath, httpHandler.New(purchasingService.UpdatePurchaseSetting))
	mux.Method("POST", GetPurchaseSettingPath, httpHandler.New(purchasingService.GetPurchaseSetting))

	mux.Method("POST", GetCheckPurchaseOrdersPath, httpHandler.New(purchasingService.GetCheckPurchaseOrders))

	mux.Method("POST", UpdatePurchaseOrderStatusPath, httpHandler.New(purchasingService.UpdatePurchaseOrderStatus))

	mux.Method("POST", UpsertReceiptOrderPath, httpHandler.New(purchasingService.UpsertReceiptOrder))
	mux.Method("POST", GetReceiptOrdersPath, httpHandler.New(purchasingService.GetReceiptOrders))
	mux.Method("POST", UpdateReceiptOrderItemsPath, httpHandler.New(purchasingService.UpdateReceiptOrderItems))
	mux.Method("POST", GetReceiptOrderItemsPath, httpHandler.New(purchasingService.GetReceiptOrderItems))
	mux.Method("POST", DeleteReceiptOrderPath, httpHandler.New(purchasingService.DeleteReceiptOrder))
	mux.Method("POST", UpdateReceiptOrderStatusPath, httpHandler.New(purchasingService.UpdateReceiptOrderStatus))

	mux.Method("POST", UpsertPurchaseInvoicePath, httpHandler.New(purchasingService.UpsertPurchaseInvoice))
	mux.Method("POST", GetPurchaseInvoicesPath, httpHandler.New(purchasingService.GetPurchaseInvoices))
	mux.Method("POST", GetPurchaseInvoiceItemsPath, httpHandler.New(purchasingService.GetPurchaseInvoiceItems))

	return mux
}
