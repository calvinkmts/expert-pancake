package impl

import (
	"context"
	"net/http"

	"github.com/calvinkmts/expert-pancake/engine/errors"
	"github.com/calvinkmts/expert-pancake/engine/httpHandler"
	db "github.com/expert-pancake/service/account/db/transaction"
	"github.com/expert-pancake/service/account/model"
)

func (a accountService) UpdateUser(w http.ResponseWriter, r *http.Request) error {

	var req model.UpdateUserRequest
	httpHandler.ParseHTTPRequest(r, &req)

	errMapRequest := a.validator.Validate(req)
	if errMapRequest != nil {
		return errors.NewClientError().WithDataMap(errMapRequest)
	}

	_, err := a.dbTrx.GetUser(context.Background(), req.AccountId)
	if err != nil {
		return errors.NewServerError(model.GetUserError, err.Error())
	}

	result, err := a.dbTrx.UpdateUserTrx(context.Background(), db.UpdateUserTrxParams{
		AccountId:   req.AccountId,
		ImageUrl:    req.ImageUrl,
		FullName:    req.FullName,
		Nickname:    req.Nickname,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Location: model.Location{
			Province:    req.Location.Province,
			Regency:     req.Location.Regency,
			District:    req.Location.District,
			FullAddress: req.Location.FullAddress,
		},
	})
	if err != nil {
		return errors.NewServerError(model.UpdateUserError, err.Error())
	}

	res := model.UpdateUserResponse{
		AccountId:   result.AccountId,
		ImageUrl:    result.ImageUrl,
		FullName:    result.FullName,
		Nickname:    result.Nickname,
		Email:       result.Email,
		PhoneNumber: result.PhoneNumber,
		Location:    result.Location,
	}
	httpHandler.WriteResponse(w, res)

	return nil
}
