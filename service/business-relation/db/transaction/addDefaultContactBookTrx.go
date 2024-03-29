package db

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	db "github.com/expert-pancake/service/business-relation/db/sqlc"
	uuid "github.com/satori/go.uuid"
)

type AddDefaultContactBookTrxParams struct {
	CompanyId   string
	CompanyName string
}

func (trx *Trx) AddDefaultContactBookTrx(ctx context.Context, arg AddDefaultContactBookTrxParams) error {

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
		var konekinId = strings.ReplaceAll(nonAlphanumericRegex.ReplaceAllString(arg.CompanyName, ""), " ", "_")

		countKonekinId, err := q.GetCountKonekinId(ctx, konekinId+"%")
		if err != nil {
			return err
		}

		if countKonekinId > 0 {
			konekinId += strconv.FormatInt(countKonekinId, 10)
		}

		id := uuid.NewV4().String()

		_, err = q.InsertContactBook(ctx, db.InsertContactBookParams{
			ID:               id,
			KonekinID:        konekinId,
			PrimaryCompanyID: arg.CompanyId,
			Name:             arg.CompanyName,
			IsDefault:        true,
		})
		if err != nil {
			return err
		}

		err = q.InsertContactBookAdditionalInfo(ctx, db.InsertContactBookAdditionalInfoParams{
			ContactBookID: id,
			Nickname:      "",
			Tag:           "",
			Note:          "",
		})
		if err != nil {
			return err
		}

		err = q.InsertContactBookMailingAddress(ctx, db.InsertContactBookMailingAddressParams{
			ContactBookID: id,
			Province:      "",
			Regency:       "",
			District:      "",
			PostalCode:    "",
			FullAddress:   "",
		})
		if err != nil {
			return err
		}

		err = q.InsertContactBookShippingAddress(ctx, db.InsertContactBookShippingAddressParams{
			ContactBookID: id,
			Province:      "",
			Regency:       "",
			District:      "",
			PostalCode:    "",
			FullAddress:   "",
		})
		if err != nil {
			return err
		}

		_, err = q.InsertContactBook(ctx, db.InsertContactBookParams{
			ID:               uuid.NewV4().String(),
			PrimaryCompanyID: arg.CompanyId,
			Name:             "UMUM",
			IsCustomer:       true,
			IsAllBranches:    true,
			IsDefault:        true,
		})
		if err != nil {
			return err
		}

		_, err = q.InsertContactBook(ctx, db.InsertContactBookParams{
			ID:               uuid.NewV4().String(),
			PrimaryCompanyID: arg.CompanyId,
			Name:             "UMUM",
			IsSupplier:       true,
			IsAllBranches:    true,
			IsDefault:        true,
		})
		if err != nil {
			return err
		}

		return err
	})

	return err
}
