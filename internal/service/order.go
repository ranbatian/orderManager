package service

import (
	"github.com/simon/orderManager/internal/db"
	"github.com/simon/orderManager/server/common"
)

func UpdateOrderType(params common.OrderType) error {
	ot, err := db.GetOrderTypeById(params.Id)
	if err != nil {
		return err
	}
	ot.OrderKindId = params.OrderKindId
	if params.Description != "" {
		ot.Description = params.Description
	}
	if params.Name != "" {
		ot.Name = params.Name
	}
	err = db.UpdateOrderType(ot)
	return err

}
