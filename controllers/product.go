package controllers

import (
	"go-api-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	// TODO: Database connection
	// TODO: Get jwt
	// TODO: Body Binding
	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	// TODO: Database connection
	// TODO: Get jwt
	// TODO: Body Binding

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title:       Product.Title,
		Description: Product.Description,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}
