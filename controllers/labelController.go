package controllers

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
	"github.com/gin-gonic/gin"
)

var bodyLabel struct {
	Label string
}

func LabelCreate(c *gin.Context) {
	c.Bind(&bodyLabel)

	label := models.Label{Label: bodyLabel.Label}
	result := initializers.DB.Create(&label)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, label)
}

func LabelIndex(c *gin.Context) {
	var labels []models.Label
	result := initializers.DB.Find(&labels)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, labels)
}

func LabelShow(c *gin.Context) {
	id := c.Param("id")

	var label models.Label
	result := initializers.DB.First(&label, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, label)
}

func LabelUpdate(c *gin.Context) {
	id := c.Param("id")

	c.Bind(&bodyLabel)

	var label models.Label
	result := initializers.DB.First(&label, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	result = initializers.DB.Model(&label).Updates(models.Label{Label: bodyLabel.Label})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, label)
}

func LabelDelete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Label{}, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}
