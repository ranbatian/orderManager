package common

import "github.com/gin-gonic/gin"

func ErrorWithMsg(ctx *gin.Context, msg string, code int) {
	ctx.JSON(200, Resp[interface{}]{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func SuccessRespond(ctx *gin.Context, data ...interface{}) {

	if len(data) > 0 {
		ctx.JSON(200, Resp[interface{}]{
			Code:    200,
			Message: "success",
			Data:    data[0],
		})
		return
	}
	ctx.JSON(200, Resp[interface{}]{
		Code:    200,
		Message: "success",
		Data:    nil,
	})
}
