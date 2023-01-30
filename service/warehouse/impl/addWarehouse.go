package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/warehouse/db/transaction"
	"github.com/expert-pancake/service/warehouse/model"
)

func (a warehouseService) AddWarehouse(w http.ResponseWriter, r *http.Request) error {

	var req model.AddWarehouseRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	arg := db.AddWarehouseTrxParams{
		BranchId: req.BranchId,
		Code:     req.Code,
		Name:     req.Name,
		Type:     req.Type,
		Address:  req.Address,
		RackIds:  req.RackIds,
	}
	result, err := a.dbTrx.AddWarehouseTrx(context.Background(), arg)

	if err != nil {
		return errors.NewServerError(model.AddWarehouseError, err.Error())
	}

	res := result
	httpHandler.WriteResponse(w, res)

	return nil
}
