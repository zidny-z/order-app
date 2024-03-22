package router

import (
	"fmt"
	"order-app/config"
	"order-app/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(c controller.Controller) error {
	port := config.GetConfigPort()
	serverInfo := fmt.Sprintf("localhost:%s", port)

	r := gin.Default()

	r.POST("/orders", c.CreateOrder)
	r.GET("/orders", c.GetOrders)
	r.PUT("/orders/:orderId", c.UpdateOrder)
	r.DELETE("/orders/:orderId", c.DeleteOrder)

	return r.Run(serverInfo)
}