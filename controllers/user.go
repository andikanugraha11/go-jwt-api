package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegistration(c *gin.Context) {
	// TODO: Database connection
	// TODO: Body Binding
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})
}

func UserLogin(c *gin.Context) {
	// TODO: Database connection
	// TODO: Body Binding

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "Invalid email/password",
		})
		return
	}

	// TODO: Compare password

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorize",
			"message": "Invalid email/password",
		})
		return
	}

	// TODO: Generate token

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
