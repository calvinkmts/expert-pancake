package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/sales/db/sqlc"
	"github.com/expert-pancake/service/sales/impl/client"
	"github.com/expert-pancake/service/sales/model"
	"github.com/expert-pancake/service/sales/util"
)

func (a salesService) GetTopCustomers(w http.ResponseWriter, r *http.Request) error {

	var req model.GetTopCustomersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetTopCustomers(context.Background(), db.GetTopCustomersParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetTopCustomersError, err.Error())
	}

	var topCustomers = make([]model.TopCustomer, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: req.CompanyId,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetTopCustomersError, err.Error())
		}
		customerName := ""
		if len(contactBook.Result) > 0 {
			customerName = contactBook.Result[0].Name
		}

		argItemVariant := client.GetItemVariantsRequest{
			Id: d.PrimaryItemVariantID,
		}
		itemVariant, err := client.GetItemVariants(argItemVariant)
		if err != nil {
			return errors.NewServerError(model.GetTopCustomersError, err.Error())
		}

		var topCustomer = model.TopCustomer{
			ContactBookId:   d.ContactBookID,
			CustomerName:    customerName,
			KonekinId:       d.KonekinID,
			ItemCode:        itemVariant.Result.ItemVariants[0].Code,
			ItemName:        itemVariant.Result.ItemVariants[0].Name,
			ItemVariantName: itemVariant.Result.ItemVariants[0].VariantName,
			Amount:           strconv.FormatInt(d.Amount, 10),
		}
		topCustomers = append(topCustomers, topCustomer)
	}

	res := model.GetTopCustomersResponse{
		TopCustomers: topCustomers,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
