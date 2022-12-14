package impl

import (
	"github.com/calvinkmts/expert-pancake/engine/validator"
	db "github.com/expert-pancake/service/business/db/transaction"
	"github.com/expert-pancake/service/business/model"
	"github.com/expert-pancake/service/business/util"
)

type businessService struct {
	config    util.Config
	validator validator.Validator
	dbTrx     db.BusinessTrx
}

func NewBusinessService(config util.Config, validator validator.Validator, dbTrx db.BusinessTrx) model.BusinessService {
	return &businessService{
		config:    config,
		validator: validator,
		dbTrx:     dbTrx,
	}
}
