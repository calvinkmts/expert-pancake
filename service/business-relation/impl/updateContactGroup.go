package impl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	"github.com/expert-pancake/service/business-relation/model"
)

func (a businessRelationService) UpdateContactGroup(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateContactGroupRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	result, err := a.dbTrx.UpdateContactGroup(context.Background(), db.UpdateContactGroupParams{
		ID:          req.GroupId,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateContactGroupError, err.Error())
	}

	members, err := a.dbTrx.GetContactBooks(context.Background(), db.GetContactBooksParams{
		PrimaryCompanyID: result.CompanyID,
		IsFilterGroupID:  true,
		ContactGroupID:   result.ID,
	})
	if err != nil {
		return errors.NewServerError(model.UpdateContactGroupError, err.Error())
	}

	res := model.UpdateContactGroupResponse{
		ContactGroupWithMember: model.ContactGroupWithMember{
			GroupId:     result.ID,
			CompanyId:   result.CompanyID,
			ImageUrl:    result.ImageUrl,
			Name:        result.Name,
			Description: result.Description,
			Member:      strconv.Itoa(int(len(members))),
		},
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
