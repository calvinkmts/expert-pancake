package db

import (
	"context"

	db "github.com/expert-pancake/service/warehouse/db/sqlc"
	"github.com/expert-pancake/service/warehouse/model"
	uuid "github.com/satori/go.uuid"
)

type AddWarehouseTrxParams struct {
	BranchId string  
	Code     string  
	Name     string  
	Type     string  
	Address  string  
	RackIds  []string
}

type AddWarehouseTrxResult struct {
	WarehouseId string
	BranchId    string
	Code        string
	Name        string
	Type        string
	Address     string
	Racks       []model.Rack 
}

func (trx *Trx) AddWarehouseTrx(ctx context.Context, arg AddWarehouseTrxParams) (AddWarehouseTrxResult, error) {
	var result AddWarehouseTrxResult

	err := trx.execTx(ctx, func(q *db.Queries) error {
		var err error

		id := uuid.NewV4().String()

		warehouseRes, err := q.AddWarehouse(ctx, db.AddWarehouseParams{
			ID:       id,
			BranchID: arg.BranchId,
			Code:     arg.Code,
			Name:     arg.Name,
			Type:     arg.Type,
			Address:  arg.Address,
		})
		if err != nil {
			return err
		}

		var racks = make([]model.Rack, 0)

		for _, rackId := range arg.RackIds {
			warehouseRack, err := q.AddWarehouseRack(ctx, db.AddWarehouseRackParams{
				WarehouseID: id,
				RackID:      rackId,
			})
			if err != nil {
				return err
			}

			rackRes, err := q.GetRack(ctx, warehouseRack.RackID)
			if err != nil {
				return err
			}

			racks = append(racks, model.Rack{
				RackId:   rackRes.ID,
				BranchId: rackRes.BranchID,
				Name:     rackRes.Name,
			})
		}

		result.WarehouseId = id
		result.BranchId = warehouseRes.BranchID
		result.Name = warehouseRes.Name
		result.Type = warehouseRes.Type
		result.Address = warehouseRes.Address
		result.Racks = racks

		return err
	})

	return result, err
}
