package controllers

import (
	"awesomeProject/models"
	"net/http"
	models "pokergame/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /user
// GET all users

func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []models.User
	db.Find(*&users)

	c.JSON(http.StatusOK, gin.H{"data": users})

}

// POST /Users
// Create new user

func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//Validate input
	var input models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {

	}
}
