package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/simon/orderManager/internal/bootstrap"
	logger "github.com/simon/orderManager/internal/log"
	"github.com/simon/orderManager/internal/op"
	"github.com/simon/orderManager/server/middleware"

	ginzap "github.com/gin-contrib/zap"
)

func InitServer() {
	bootstrap.Init()
	r := gin.New()

	r.Use(ginzap.Ginzap(logger.Log, time.RFC3339, true))

	r.Use(ginzap.RecoveryWithZap(logger.Log, true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.POST("/login", op.Login)
	rgv1 := r.Group("/v1")
	rgv1.Use(middleware.JwtMiddleware())
	rgUser := rgv1.Group("/user")
	rgOrder := rgv1.Group("/order")
	order(rgOrder)
	user(rgUser)
	r.Run(fmt.Sprintf("0.0.0.0:%d", 8080)) // 监
}

func user(g *gin.RouterGroup) {
	g.GET("/", op.CurrentUser)
	g.POST("/", op.CreateUser)
}

func order(g *gin.RouterGroup) {
	g.POST("/type", op.CreateOrderType)
	g.PUT("/type", op.UpdateOrderType)
	g.POST("/kind", op.CreateOrderKind)
	g.POST("/kind/:id", op.CreateOrderKind)
}
