package controller

import (
	"net/http"
	"order-app/database"
	"order-app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func (c Controller) CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error binding JSON request",
		})
		return
	}

	orderResult, err := c.db.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error creating order",
		})
		return
	}

	ctx.JSON(http.StatusCreated, orderResult)
}

func (c Controller) GetOrders(ctx *gin.Context) {
	orders, err := c.db.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error getting data",
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c Controller) UpdateOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error binding JSON request",
		})
		return
	}

	orderResult, err, isFound := c.db.UpdateOrder(orderID, newOrder)
	if err != nil {
		if !isFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    "404",
				"message": err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error updating order",
		})
		return
	}

	ctx.JSON(http.StatusOK, orderResult)
}

func (c Controller) DeleteOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Invalid param orderId",
		})
		return
	}

	err, isFound := c.db.DeleteOrder(orderID)
	if err != nil {
		if !isFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    "404",
				"message": err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "500",
			"message": "Error deleting order",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "Successfully deleted order",
	})
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}
