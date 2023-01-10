package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/transaction"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) UpdateCustomer(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCustomerRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	credit_limit, _ := strconv.ParseInt(req.CreditLimit, 10, 64)
	payment_term, _ := strconv.ParseInt(req.PaymentTerm, 10, 64)
	arg := db.UpdateCustomerTrxParams{
		Id:          req.ContactBookId,
		IsTax:       req.IsTax,
		TaxId:       req.TaxId,
		Pic:         req.Pic,
		CreditLimit: credit_limit,
		PaymentTerm: int32(payment_term),
	}

	result, err := a.dbTrx.UpdateCustomerTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateCustomerError, err.Error())
	}

	res := model.UpdateCustomerResponse{
		CustomerInfo: model.CustomerInfo{
			ContactBookId:    result.ContactBookId,
			ContactGroupName: result.ContactGroupName,
			Name:             result.Name,
			Email:            result.Email,
			Phone:            result.Phone,
			Mobile:           result.Mobile,
			Web:              result.Web,
			IsTax:            result.IsTax,
			TaxId:            result.TaxId,
			Pic:              result.Pic,
			CreditLimit:      result.CreditLimit,
			PaymentTerm:      result.PaymentTerm,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
