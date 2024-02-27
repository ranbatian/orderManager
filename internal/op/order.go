package op

import (
	"github.com/gin-gonic/gin"
	"github.com/simon/orderManager/internal/db"
	logger "github.com/simon/orderManager/internal/log"
	"github.com/simon/orderManager/internal/model"
	"github.com/simon/orderManager/internal/service"
	"github.com/simon/orderManager/server/common"
)

func CreateOrderType(c *gin.Context) {
	var odType model.OrderType
	if err := c.ShouldBind(&odType); err == nil {
		err := db.CreateOrderType(odType)
		if err != nil {
			logger.Log.Sugar().Error("create order type error. %s", odType.Name)
			common.ErrorWithMsg(c, "创建类型失败", 500)
			return
		}
		common.SuccessRespond(c)
		return
	}
	common.ErrorWithMsg(c, "参数传递错误", 500)
}

func UpdateOrderType(c *gin.Context) {
	var odType common.OrderType
	if err := c.ShouldBind(&odType); err == nil {
		err = service.UpdateOrderType(odType)
		if err != nil {
			logger.Log.Sugar().Error("create order type error. %s", odType.Name)
			common.ErrorWithMsg(c, "编辑类型失败", 500)
			return
		}
		common.SuccessRespond(c)
		return
	}
	common.ErrorWithMsg(c, "参数传递错误", 500)
}
func CreateOrderKind(c *gin.Context) {
	var odType model.OrderKind
	if err := c.ShouldBind(&odType); err == nil {
		err := db.CreateOrderKind(odType)
		if err != nil {
			logger.Log.Sugar().Error("create order kind error. %s", odType.Name)
			common.ErrorWithMsg(c, "创建类目失败", 500)
			return
		}
		common.SuccessRespond(c)
		return
	}
	common.ErrorWithMsg(c, "参数传递错误", 500)
	return
}

func CreateOrder(c *gin.Context) {
	var order model.Order
	userId := c.GetUint("userId")
	if err := c.ShouldBind(&order); err == nil {
		err := service.CreateOrder(userId, order)
		if err != nil {
			logger.Log.Sugar().Error("create order kind error. %d", order.ID)
			common.ErrorWithMsg(c, "创建账单失败", 500)
			return
		}
		common.SuccessRespond(c)
		return
	}
	common.ErrorWithMsg(c, "参数传递错误", 500)
	return
}
