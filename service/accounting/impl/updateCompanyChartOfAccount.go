package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/accounting/db/transaction"
	"github.com/expert-pancake/service/accounting/model"
)

func (a accountingService) UpdateCompanyChartOfAccount(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyChartOfAccountRequest

	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	var branches = make([]string, 0)
	if req.Branches != nil {
		branches = req.Branches
	}

	arg := db.UpdateChartOfAccountTrxParams{
		Id:                    req.ChartOfAccountId,
		CurrencyCode:          req.CurrencyCode,
		ChartOfAccountGroupId: req.ChartOfAccountGroupId,
		AccountCode:           req.AccountCode,
		AccountName:           req.AccountName,
		BankName:              req.BankName,
		BankAccountNumber:     req.BankAccountNumber,
		BankCode:              req.BankCode,
		IsAllBranches:         req.IsAllBranches,
		Branches:              branches,
		IsDeleted:             req.IsDeleted,
	}

	result, err := a.dbTrx.UpdateChartOfAccountTrx(context.Background(), arg)
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyChartOfAccountError, err.Error())
	}

	resultGroup, err := a.dbTrx.GetChartOfAccountGroup(context.Background(), arg.ChartOfAccountGroupId)
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyChartOfAccountError, err.Error())
	}

	res := model.UpsertCompanyChartOfAccountResponse{
		ChartOfAccount: model.ChartOfAccount{
			ChartOfAccountId:      result.ChartOfAccountId,
			CompanyId:             result.CompanyId,
			CurrencyCode:          result.CurrencyCode,
			ChartOfAccountGroupId: result.ChartOfAccountGroupId,
			ReportType:            resultGroup.ReportType,
			AccountType:           resultGroup.AccountType,
			AccountGroup:          resultGroup.AccountGroupName,
			AccountCode:           result.AccountCode,
			AccountName:           result.AccountName,
			BankName:              result.BankName,
			BankAccountNumber:     result.BankAccountNumber,
			BankCode:              result.BankCode,
			IsAllBranches:         result.IsAllBranches,
			Branches:              result.Branches,
			IsDeleted:             result.IsDeleted,
		},
	}

	httpHandler.WriteResponse(w, res)

	return nil
}
