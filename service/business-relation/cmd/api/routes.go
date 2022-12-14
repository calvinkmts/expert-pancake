package main

import (
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	"github.com/expert-pancake/service/business-relation/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	AddContactGroupPath    = "/business-relation/contact-group/add"
	UpdateContactGroupPath = "/business-relation/contact-group/update"
	GetContactGroupsPath   = "/business-relation/contact-groups"
	AddContactBookPath     = "/business-relation/contact-book/add"
	UpdateContactBookPath  = "/business-relation/contact-book/update"
	GetContactBooksPath    = "/business-relation/contact-books"
	UpdateCustomerPath     = "/business-relation/customer/update"
	GetCustomersPath       = "/business-relation/customers"
	UpdateSupplierPath     = "/business-relation/supplier/update"
	GetSuppliersPath       = "/business-relation/suppliers"
)

func (c *component) Routes(businessRelationService model.BusinessRelationService) http.Handler {
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

	mux.Method("POST", AddContactGroupPath, httpHandler.New(businessRelationService.AddContactGroup))
	mux.Method("POST", UpdateContactGroupPath, httpHandler.New(businessRelationService.UpdateContactGroup))
	mux.Method("POST", GetContactGroupsPath, httpHandler.New(businessRelationService.GetContactGroups))
	mux.Method("POST", AddContactBookPath, httpHandler.New(businessRelationService.AddContactBook))
	mux.Method("POST", UpdateContactBookPath, httpHandler.New(businessRelationService.UpdateContactBook))
	mux.Method("POST", GetContactBooksPath, httpHandler.New(businessRelationService.GetContactBooks))
	mux.Method("POST", UpdateCustomerPath, httpHandler.New(businessRelationService.UpdateCustomer))
	mux.Method("POST", GetCustomersPath, httpHandler.New(businessRelationService.GetCustomers))
	mux.Method("POST", UpdateSupplierPath, httpHandler.New(businessRelationService.UpdateSupplier))
	mux.Method("POST", GetSuppliersPath, httpHandler.New(businessRelationService.GetSuppliers))

	return mux
}
