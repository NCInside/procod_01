package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodySubm struct {
	Code        string
	Runtime     uint
	Is_correct  bool
	Error       string
	Memory      uint
	UserID      uint
	ChallengeID uint
}

var body struct {
	UserID      uint `form:"userid"`
	ChallengeID uint `form:"challengeid"`
}

func SubmissionCreate(c *gin.Context) {
	c.Bind(&bodySubm)

	submission := models.Submission{
		Code:        bodySubm.Code,
		Runtime:     bodySubm.Runtime,
		Is_correct:  bodySubm.Is_correct,
		Error:       bodySubm.Error,
		Memory:      bodySubm.Memory,
		UserID:      bodySubm.UserID,
		ChallengeID: bodySubm.ChallengeID,
	}
	result := initializers.DB.Create(&submission)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, submission)
}

func SubmissionIndex(c *gin.Context) {
	var submissions []models.Submission
	initializers.DB.Find(&submissions)

	c.JSON(200, submissions)
}

func SubmissionShow(c *gin.Context) {
	id := c.Param("id")

	var submission models.Submission
	initializers.DB.First(&submission, id)

	c.JSON(200, submission)
}

func SubmissionUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodySubm)

	var submission models.Submission
	result := initializers.DB.First(&submission, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	result = initializers.DB.Model(&submission).Updates(models.Submission{
		Code:        bodySubm.Code,
		Runtime:     bodySubm.Runtime,
		Is_correct:  bodySubm.Is_correct,
		Error:       bodySubm.Error,
		Memory:      bodySubm.Memory,
		UserID:      bodySubm.UserID,
		ChallengeID: bodySubm.ChallengeID,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, submission)
}

func SubmissionDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Submission{}, id)

	c.Status(200)
}

func SubmissionShowIndex(c *gin.Context) {
	c.Bind(&body)

	if body.ChallengeID == 0 && body.UserID == 0 {
		c.JSON(400, "Parameter not found")

	} else {
		var submissions []models.Submission

		if body.ChallengeID == 0 {
			initializers.DB.Where("user_id = ?", body.UserID).Find(&submissions)

		} else if body.UserID == 0 {
			initializers.DB.Where("challenge_id = ?", body.ChallengeID).Find(&submissions)

		} else {
			initializers.DB.Where("user_id = ? AND challenge_id = ?", body.UserID, body.ChallengeID).Find(&submissions)

		}

		c.JSON(200, submissions)
	}
}
