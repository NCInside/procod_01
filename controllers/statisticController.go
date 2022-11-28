package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyStat struct {
	Total_memory            uint
	Total_runtime           uint
	Num_challenge_attempted uint
	Num_challenge_completed uint
	Num_challenge_made      uint
	UserID                  uint
}

func StatisticCreate(c *gin.Context) {
	c.Bind(&bodyStat)

	statistic := models.Statistic{
		Total_memory:            bodyStat.Total_memory,
		Total_runtime:           bodyStat.Total_runtime,
		Num_challenge_attempted: bodyStat.Num_challenge_attempted,
		Num_challenge_completed: bodyStat.Num_challenge_completed,
		Num_challenge_made:      bodyStat.Num_challenge_made,
		UserID:                  bodyStat.UserID,
	}
	result := initializers.DB.Create(&statistic)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, statistic)
}

func StatisticIndex(c *gin.Context) {
	var statistics []models.Statistic
	initializers.DB.Find(&statistics)

	c.JSON(200, statistics)
}

func StatisticShow(c *gin.Context) {
	id := c.Param("id")

	var statistic models.Statistic
	initializers.DB.First(&statistic, id)

	c.JSON(200, statistic)
}

func StatisticUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyStat)

	var statistic models.Statistic
	initializers.DB.First(&statistic, id)

	initializers.DB.Model(&statistic).Updates(models.Statistic{
		Total_memory:            bodyStat.Total_memory,
		Total_runtime:           bodyStat.Total_runtime,
		Num_challenge_attempted: bodyStat.Num_challenge_attempted,
		Num_challenge_completed: bodyStat.Num_challenge_completed,
		Num_challenge_made:      bodyStat.Num_challenge_made,
		UserID:                  bodyStat.UserID,
	})

	c.JSON(200, statistic)
}

func StatisticDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Statistic{}, id)

	c.Status(200)
}

func StatisticIndexUser(c *gin.Context) {
	userID := c.Param("userid")

	var statistics []models.Statistic
	initializers.DB.Where("user_id = ?", userID).Find(&statistics)

	c.JSON(200, statistics)
}
