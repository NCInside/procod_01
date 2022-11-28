package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyUser struct {
	Username string
	Password string
}

func UserCreate(c *gin.Context) {
	c.Bind(&bodyUser)

	user := models.User{Username: bodyUser.Username, Password: bodyUser.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, user)
}

func UserIndex(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, users)
}

func UserShow(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, user)
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyUser)

	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Username: bodyUser.Username,
		Password: bodyUser.Password,
	})

	c.JSON(200, user)
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.User{}, id)

	c.Status(200)
}