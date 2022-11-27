package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/sqlc"
	"github.com/expert-pancake/service/accounting/model"
	uuid "github.com/satori/go.uuid"
)

func (a accountingService) SetCompanyFirstBankAccount(w http.ResponseWriter, r *http.Request) error {

	var req model.SetCompanyFirstBankAccountRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	id := uuid.NewV4().String()

	arg := db.UpsertCompanyChartOfAccountParams{
		ID:                id,
		CompanyID:         req.CompanyId,
		BranchID:          req.BranchId,
		AccountCode:       "10000",
		AccountName:       "BANK",
		AccountGroup:      "BANK",
		BankName:          req.BankName,
		BankAccountNumber: req.BankAccountNumber,
		BankCode:          req.BankCode,
		OpeningBalance:    0,
		IsDeleted:         0,
	}

	result, err := a.dbTrx.UpsertCompanyChartOfAccount(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.SetCompanyFirstBankAccountError, err.Error())
	}

	res := model.UpsertCompanyChartOfAccountResponse{
		ChartOfAccount: model.ChartOfAccount{
			ChartOfAccountId:  result.ID,
			CompanyId:         result.CompanyID,
			BranchId:          result.BranchID,
			AccountCode:       result.AccountCode,
			AccountName:       result.AccountName,
			AccountGroup:      result.AccountGroup,
			BankName:          result.BankName,
			BankAccountNumber: result.BankAccountNumber,
			BankCode:          result.BankCode,
			OpeningBalance:    result.OpeningBalance,
			IsDeleted:         result.IsDeleted,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}