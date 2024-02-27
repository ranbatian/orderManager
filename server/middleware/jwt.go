package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	logger "github.com/simon/orderManager/internal/log"
	"github.com/simon/orderManager/pkg/util"
	"github.com/simon/orderManager/server/common"
)

func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("authorization")
		if authorization == "" {
			c.JSON(200, common.Resp[interface{}]{
				Code:    403,
				Message: "请求头中 authorization 为空",
				Data:    nil,
			})
			c.Abort()
			return
		}
		// parts := strings.Split(authorization, ",")
		// if len(parts) != 3 {
		// 	c.JSON(200, common.Resp[interface{}]{
		// 		Code:    403,
		// 		Message: "请求头中 authorization 格式有误",
		// 		Data:    nil,
		// 	})
		// 	c.Abort()
		// 	return
		// }

		s := strings.Split(authorization, " ")
		token := s[1]
		mc, err := util.ParseToken(token)
		if err != nil {
			logger.Log.Sugar().Error("ParseToken error. %s", err.Error())
			c.JSON(200, common.Resp[interface{}]{
				Code:    401,
				Message: "解析 token 失败.",
				Data:    nil,
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Set("userId", mc.UserId)
		c.Next()

	}
}
