package controllers

import (
	"strconv"

	"github.com/NCInside/procod_01/initializers"
	"github.com/gin-gonic/gin"
)

func ChallengeLabelCreate(c *gin.Context) {

	challenge := c.Query("challengeid")
	label := c.Query("labelid")

	challengeId, _ := strconv.Atoi(challenge)
	labelId, _ := strconv.Atoi(label)

	result := initializers.DB.Table("challenge_labels").Create(map[string]interface{}{
		"challenge_id": challengeId, "label_id": labelId,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.Status(200)

}

func ChallengeLabelDelete(c *gin.Context) {

	challenge := c.Query("challengeid")
	label := c.Query("labelid")

	challengeId, _ := strconv.Atoi(challenge)
	labelId, _ := strconv.Atoi(label)

	result := initializers.DB.Unscoped().
		Exec("DELETE FROM `challenge_labels` WHERE challenge_id = ? AND label_id = ?", challengeId, labelId)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.Status(200)

}
