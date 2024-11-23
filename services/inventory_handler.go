package services

import (
	"fmt"
	"go-asg4/config"
	"go-asg4/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadProductStock(c *gin.Context) {
	productID := c.Query("product_id")

	if productID == "" {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("product ID is required"))
		return
	}

	var inventory models.Inventory
	query := config.DB.Table("inventories").
		Select("inventories.stock").
		Joins("JOIN products ON products.product_id = inventories.product_id")

	if productID != "" {
		query = query.Where("inventories.product_id = ?", productID)
	}
	err := query.First(&inventory).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse("no matching inventory found"))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to fetch inventory: %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse("stock retrieved successfully", inventory.Stock))
}

func UpdateInventory(c *gin.Context) {
	productID := c.Param("product_id")
	if productID == "" {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("product ID is required"))
		return
	}

	var stockUpdate struct {
		StockAdd int `json:"stock_add"`
	}
	if err := c.ShouldBindJSON(&stockUpdate); err != nil {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())))
		return
	}

	var inventory models.Inventory
	err := config.DB.First(&inventory, "product_id = ?", productID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("no inventory found for product ID %s", productID)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve inventory: %s", err.Error())))
		return
	}

	inventory.Stock += stockUpdate.StockAdd
	if err := config.DB.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to update inventory: %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse("inventory updated successfully", inventory))
}
