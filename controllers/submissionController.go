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
	initializers.DB.First(&submission, id)

	initializers.DB.Model(&submission).Updates(models.Submission{
		Code:        bodySubm.Code,
		Runtime:     bodySubm.Runtime,
		Is_correct:  bodySubm.Is_correct,
		Error:       bodySubm.Error,
		Memory:      bodySubm.Memory,
		UserID:      bodySubm.UserID,
		ChallengeID: bodySubm.ChallengeID,
	})

	c.JSON(200, submission)
}

func SubmissionDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Submission{}, id)

	c.Status(200)
}
