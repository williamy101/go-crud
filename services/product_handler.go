package services

import (
	"fmt"
	"go-asg4/config"
	"go-asg4/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			models.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())))
		return
	}
	err = config.DB.Create(&product).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.NewFailedResponse(fmt.Sprintf("failed to save product: %s", err.Error())))
		return
	}

	fmt.Printf("Product created: %+v\n", product)
	c.JSON(http.StatusCreated, models.NewSuccessResponse("product created successfully", product))
}

func ReadProducts(c *gin.Context) {
	var products []models.Product

	err := config.DB.Find(&products).Error
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			models.NewFailedResponse(fmt.Sprintf("failed to fetch products: %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, models.NewSuccessResponse("products retrieved successfully", products))
}

func ReadProductByIdOrCategory(c *gin.Context) {
	var products []models.Product
	id := c.Query("id")
	category := c.Query("category")
	if id == "" && category == "" {
		c.JSON(
			http.StatusBadRequest,
			models.NewFailedResponse("product ID or category is required"))
		return
	}

	query := config.DB
	if id != "" {
		query = query.Where("product_id = ?", id)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Find(&products).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(
				http.StatusNotFound,
				models.NewFailedResponse("no products matching the criteria"))
			return
		}
		c.JSON(
			http.StatusInternalServerError,
			models.NewFailedResponse(fmt.Sprintf("failed to fetch product: %s", err.Error())))
		return
	}

	if len(products) == 0 {
		c.JSON(
			http.StatusNotFound,
			models.NewFailedResponse("no products found"))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse("products retrieved successfully", products))
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			models.NewFailedResponse("product ID required"))
		return
	}

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse(fmt.Sprintf("failed to bind request: %s", err.Error())))
		return
	}

	var existingProduct models.Product
	if err := config.DB.First(&existingProduct, "product_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("product with ID %s not found", id)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve product: %s", err.Error())))
		return
	}

	err := config.DB.Model(&existingProduct).Updates(product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse("product updated successfully", existingProduct))
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("product ID is required"))
		return
	}

	var product models.Product
	if err := config.DB.First(&product, "product_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("product with ID %s not found", id)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve product: %s", err.Error())))
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to delete product: %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse(fmt.Sprintf("product with ID %s deleted successfully", id), nil))
}

func UploadProductImage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("product ID is required"))
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving file:", err.Error())
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("failed to retrieve image"))
		return
	}

	fmt.Printf("File received: %+v\n", file.Filename)

	filePath := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to save image: %s", err.Error())))
		return
	}

	var product models.Product
	if err := config.DB.First(&product, "product_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("product with ID %s not found", id)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve product: %s", err.Error())))
		return
	}

	product.ImagePath = &filePath
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to update product: %s", err.Error())))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse("image uploaded successfully", gin.H{"image_path": filePath}))
}

func DownloadProductImage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.NewFailedResponse("product ID is required"))
		return
	}

	var product models.Product
	if err := config.DB.First(&product, "product_id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.NewFailedResponse(fmt.Sprintf("product with ID %s not found", id)))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewFailedResponse(fmt.Sprintf("failed to retrieve product: %s", err.Error())))
		return
	}

	if product.ImagePath == nil || *product.ImagePath == "" {
		c.JSON(http.StatusNotFound, models.NewFailedResponse("no image found for this product"))
		return
	}
	filePath := *product.ImagePath
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, models.NewFailedResponse("image file not found on the server"))
		return
	}

	c.File(filePath)
}
