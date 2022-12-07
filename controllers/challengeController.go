package controllers

import (
	"strconv"

	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	label := c.Query("label")

	var challenges []models.Challenge
	var result *gorm.DB

	labelId, err := strconv.Atoi(label)

	if err != nil {

		result = initializers.DB.
			Preload("ChallengeLabels").Preload("ChallengeExamples").Preload("ChallengeTargets").Find(&challenges)

	} else {

		result = initializers.DB.Where(
			"EXISTS (?)",
			initializers.DB.Table("challenge_labels").Where(
				"challenges.id = challenge_labels.challenge_id AND challenge_labels.label_id = ?", labelId),
		).Preload("ChallengeLabels").Preload("ChallengeExamples").Preload("ChallengeTargets").Find(&challenges)

	}

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challenges)
}

func ChallengeShow(c *gin.Context) {
	id := c.Param("id")

	var challenge models.Challenge
	result := initializers.DB.Preload("ChallengeLabels").Preload("ChallengeExamples").Preload("ChallengeTargets").First(&challenge, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challenge)
}

func ChallengeUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyChal)

	var challenge models.Challenge
	result := initializers.DB.First(&challenge, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	result = initializers.DB.Model(&challenge).Updates(models.Challenge{
		Title:       bodyChal.Title,
		Description: bodyChal.Description,
		UserID:      bodyChal.UserID,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

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

func ChallengeSubmission(c *gin.Context) {
	id := c.Param("id")

	var challenges []models.Challenge
	result := initializers.DB.Where("id IN (?)",
		initializers.DB.Table("submissions").Select("challenge_id").
			Where("is_correct = 0 AND user_id = ?", id)).Find(&challenges)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, challenges)

}
