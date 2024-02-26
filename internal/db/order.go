package db

import "github.com/simon/orderManager/internal/model"

func CreateOrderType(odType model.OrderType) error {
	return DB.Create(&odType).Error
}

func GetOrderTypeById(id uint) (orderTy model.OrderType, err error) {
	err = DB.Model(&model.OrderType{}).Where("id = ?", id).Association("OrderKind").Find(&orderTy)
	return
}

func UpdateOrderType(odType model.OrderType) error {
	return DB.Save(&odType).Error
}

func CreateOrderKind(odType model.OrderKind) error {
	return DB.Create(&odType).Error
}
