package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyChalEx struct {
	Ex_input    string
	Ex_output   string
	Description string
	ChallengeID uint
}

func ChallengeExCreate(c *gin.Context) {
	c.Bind(&bodyChalEx)

	challengeEx := models.ChallengeExample{
		Ex_input:    bodyChalEx.Ex_input,
		Ex_output:   bodyChalEx.Ex_output,
		Description: bodyChalEx.Description,
		ChallengeID: bodyChalEx.ChallengeID,
	}
	result := initializers.DB.Create(&challengeEx)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challengeEx)
}

func ChallengeExIndex(c *gin.Context) {
	var challengesEx []models.ChallengeExample
	initializers.DB.Find(&challengesEx)

	c.JSON(200, challengesEx)
}

func ChallengeExShow(c *gin.Context) {
	id := c.Param("id")

	var challengeEx models.ChallengeExample
	initializers.DB.First(&challengeEx, id)

	c.JSON(200, challengeEx)
}

func ChallengeExUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyChalEx)

	var challengeEx models.ChallengeExample
	result := initializers.DB.First(&challengeEx, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	result = initializers.DB.Model(&challengeEx).Updates(models.ChallengeExample{
		Ex_input:    bodyChalEx.Ex_input,
		Ex_output:   bodyChalEx.Ex_output,
		Description: bodyChalEx.Description,
		ChallengeID: bodyChalEx.ChallengeID,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challengeEx)
}

func ChallengeExDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.ChallengeExample{}, id)

	c.Status(200)
}

func ChallengeExIndexChallenge(c *gin.Context) {
	userID := c.Param("challengeid")

	var challengesEx []models.ChallengeExample
	initializers.DB.Where("challenge_id = ?", userID).Find(&challengesEx)

	c.JSON(200, challengesEx)
}
