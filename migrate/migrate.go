package main

import (
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{}, &models.Statistic{}, &models.Challenge{},
		&models.ChallengeExample{}, &models.Label{}, &models.ChallengeTarget{}, &models.Submission{})
}
