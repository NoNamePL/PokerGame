package controllers

import (
	models "awesomeProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /user
// GET all users

func FindUsers(c *gin.Context) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create User
	user := models.User{
		Name:     input.Name,
		Password: input.Password,
		Phone:    input.Password,
		Email:    input.Email,
	}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /Users/:id
// Find a user
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//GET model if exist

	var user models.User
	if err := db.Where("id = ?",
		c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}
	return
}

// Patch /Users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// valide input
	var input models.UpdateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /books/:id
// Delete a user

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// GET model if exist
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
