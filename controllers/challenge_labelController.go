package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyChalLb struct {
	Label       string
	ChallengeID uint
}

func ChallengeLbCreate(c *gin.Context) {
	c.Bind(&bodyChalLb)

	challengeLb := models.ChallengeLabel{Label: bodyChalLb.Label, ChallengeID: bodyChalLb.ChallengeID}
	result := initializers.DB.Create(&challengeLb)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challengeLb)
}

func ChallengeLbIndex(c *gin.Context) {
	var challengesLb []models.ChallengeLabel
	initializers.DB.Find(&challengesLb)

	c.JSON(200, challengesLb)
}

func ChallengeLbShow(c *gin.Context) {
	id := c.Param("id")

	var challengeLb models.ChallengeLabel
	initializers.DB.First(&challengeLb, id)

	c.JSON(200, challengeLb)
}

func ChallengeLbUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyChalLb)

	var challengeLb models.ChallengeLabel
	initializers.DB.First(&challengeLb, id)

	initializers.DB.Model(&challengeLb).Updates(models.ChallengeLabel{
		Label:       bodyChalLb.Label,
		ChallengeID: bodyChalLb.ChallengeID,
	})

	c.JSON(200, challengeLb)
}

func ChallengeLbDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.ChallengeLabel{}, id)

	c.Status(200)
}

func ChallengeLbIndexChallenge(c *gin.Context) {
	userID := c.Param("challengeid")

	var challengesLb []models.ChallengeLabel
	initializers.DB.Where("challenge_id = ?", userID).Find(&challengesLb)

	c.JSON(200, challengesLb)
}
