package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyChal struct {
	Title       string
	Description string
	UserID      uint
}

func ChallengeCreate(c *gin.Context) {
	c.Bind(&bodyChal)

	challenge := models.Challenge{Title: bodyChal.Title, Description: bodyChal.Description, UserID: bodyChal.UserID}
	result := initializers.DB.Create(&challenge)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challenge)
}

func ChallengeIndex(c *gin.Context) {
	var challenges []models.Challenge
	initializers.DB.Find(&challenges)

	c.JSON(200, challenges)
}

func ChallengeShow(c *gin.Context) {
	id := c.Param("id")

	var challenge models.Challenge
	initializers.DB.First(&challenge, id)

	c.JSON(200, challenge)
}

func ChallengeUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyChal)

	var challenge models.Challenge
	initializers.DB.First(&challenge, id)

	initializers.DB.Model(&challenge).Updates(models.Challenge{
		Title:       bodyChal.Title,
		Description: bodyChal.Description,
		UserID:      bodyChal.UserID,
	})

	c.JSON(200, challenge)
}

func ChallengeDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Challenge{}, id)

	c.Status(200)
}

func ChallengeIndexUser(c *gin.Context) {
	userID := c.Param("userid")

	var challenges []models.Challenge
	initializers.DB.Where("user_id = ?", userID).Find(&challenges)

	c.JSON(200, challenges)
}
