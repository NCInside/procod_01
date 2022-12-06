package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyChalTr struct {
	Input         string
	Target_output string
	ChallengeID   uint
}

func ChallengeTrCreate(c *gin.Context) {
	c.Bind(&bodyChalTr)

	challengeTr := models.ChallengeTarget{
		Input:         bodyChalTr.Input,
		Target_output: bodyChalTr.Target_output,
		ChallengeID:   bodyChalTr.ChallengeID,
	}
	result := initializers.DB.Create(&challengeTr)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challengeTr)
}

func ChallengeTrIndex(c *gin.Context) {
	var challengesTr []models.ChallengeTarget
	initializers.DB.Find(&challengesTr)

	c.JSON(200, challengesTr)
}

func ChallengeTrShow(c *gin.Context) {
	id := c.Param("id")

	var challengeTr models.ChallengeTarget
	initializers.DB.First(&challengeTr, id)

	c.JSON(200, challengeTr)
}

func ChallengeTrUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyChalTr)

	var challengeTr models.ChallengeTarget
	result := initializers.DB.First(&challengeTr, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	result = initializers.DB.Model(&challengeTr).Updates(models.ChallengeTarget{
		Input:         bodyChalTr.Input,
		Target_output: bodyChalTr.Target_output,
		ChallengeID:   bodyChalTr.ChallengeID,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challengeTr)
}

func ChallengeTrDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.ChallengeTarget{}, id)

	c.Status(200)
}

func ChallengeTrIndexChallenge(c *gin.Context) {
	userID := c.Param("challengeid")

	var challengesTr []models.ChallengeTarget
	initializers.DB.Where("challenge_id = ?", userID).Find(&challengesTr)

	c.JSON(200, challengesTr)
}
