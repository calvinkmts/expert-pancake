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

func (a salesService) GetDeliveryOrders(w http.ResponseWriter, r *http.Request) error {

	var req model.GetDeliveryOrdersRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.GetDeliveryOrders(context.Background(), db.GetDeliveryOrdersParams{
		CompanyID: req.CompanyId,
		BranchID:  req.BranchId,
		StartDate: util.StringToDate(req.StartDate),
		EndDate:   util.StringToDate(req.EndDate),
	})
	if err != nil {
		return errors.NewServerError(model.GetDeliveryOrdersError, err.Error())
	}

	var deliveryOrders = make([]model.DeliveryOrder, 0)

	for _, d := range result {
		argContactBook := client.GetContactBooksRequest{
			Id:        d.ContactBookID,
			CompanyId: d.CompanyID,
		}
		contactBook, err := client.GetContactBooks(argContactBook)
		if err != nil {
			return errors.NewServerError(model.GetDeliveryOrdersError, err.Error())
		}
		customerName := ""
		if len(contactBook.Result) > 0 {
			customerName = contactBook.Result[0].Name
		}

		argWarehouse := client.GetWarehousesRequest{
			Id:       d.WarehouseID,
			BranchId: "1",
		}
		warehouseName := ""
		warehouse, err := client.GetWarehouses(argWarehouse)
		if len(warehouse.Result.Warehouses) > 0 {
			warehouseName = warehouse.Result.Warehouses[0].Name
		}

		var deliveryOrder = model.DeliveryOrder{
			TransactionId:        d.ID,
			CompanyId:            d.CompanyID,
			BranchId:             d.BranchID,
			WarehouseId:          "",
			WarehouseName:        warehouseName,
			FormNumber:           d.FormNumber,
			TransactionDate:      d.TransactionDate.Format(util.DateLayoutYMD),
			ContactBookId:        d.ContactBookID,
			SecondaryCompanyId:   d.SecondaryCompanyID,
			CustomerName:         customerName,
			KonekinId:            d.KonekinID,
			SalesOrderId:         d.SalesOrderID,
			SalesOrderFormNumber: d.SalesOrderFormNumber,
			TotalItems:           strconv.FormatInt(d.TotalItems, 10),
			Status:               d.Status,
		}
		deliveryOrders = append(deliveryOrders, deliveryOrder)
	}

	res := model.GetDeliveryOrdersResponse{
		DeliveryOrders: deliveryOrders,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
