package services

import (
	"fmt"
	"go-asg4/config"
	"go-asg4/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())))
		return
	}

	var product models.Product
	if err := config.DB.First(&product, "product_id = ?", order.ProductID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("product with ID %d not found", order.ProductID)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve product: %s", err.Error())))
		return
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to create order: %s", err.Error())))
		return
	}

	c.JSON(http.StatusCreated, models.NewSuccessResponse("order created successfully", order))
}

func ReadOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("order ID is required"))
		return
	}

	var order models.Order
	err := config.DB.Preload("Product").First(&order, "order_id = ?", orderID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("order with ID %s not found", orderID)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve order: %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse("order retrieved successfully", order))
}
