package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business/db/sqlc"
	"github.com/expert-pancake/service/business/model"
)

func (a businessService) UpdateCompany(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateCompanyRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.UpdateCompany(context.Background(), db.UpdateCompanyParams{
		ID:                req.CompanyId,
		Name:              req.Name,
		InitialName:       req.InitialName,
		Type:              req.Type,
		ResponsiblePerson: req.ResponsiblePerson,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateCompanyError, err.Error())
	}

	var companyBranches = make([]model.CompanyBranch, 0)
	resultBranches, err := a.dbTrx.GetUserCompanyBranches(context.Background(), db.GetUserCompanyBranchesParams{
		UserID:    req.AccountId,
		CompanyID: req.CompanyId,
	})
	if err != nil {
		return errors.NewServerError(model.GetUserCompanyBranchesError, err.Error())
	}

	for _, d := range resultBranches {
		var companyBranch = model.CompanyBranch{
			AccountId:   d.UserID,
			CompanyId:   d.CompanyID,
			BranchId:    d.ID,
			Name:        d.Name,
			Address:     d.Address,
			PhoneNumber: d.PhoneNumber,
			IsCentral:   d.IsCentral,
		}
		companyBranches = append(companyBranches, companyBranch)
	}

	res := model.RegisterCompanyResponse{
		Company: model.Company{
			AccountId:         result.UserID,
			CompanyId:         result.ID,
			Name:              result.Name,
			InitialName:       result.InitialName,
			Type:              result.Type,
			ResponsiblePerson: result.ResponsiblePerson,
			Branches:          companyBranches,
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
