package main

import (
	"github.com/NCInside/procod_01/controllers"
	"github.com/NCInside/procod_01/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	//User
	router.POST("/users", controllers.UserCreate)
	router.PUT("/users/:id", controllers.UserUpdate)
	router.GET("/users", controllers.UserIndex)
	router.GET("/users/:id", controllers.UserShow)
	router.DELETE("/users/:id", controllers.UserDelete)

	//Statistic
	router.POST("/statistics", controllers.StatisticCreate)
	router.PUT("/statistics/:id", controllers.StatisticUpdate)
	router.GET("/statistics", controllers.StatisticIndex)
	router.GET("/statistics/:id", controllers.StatisticShow)
	router.DELETE("/statistics/:id", controllers.StatisticDelete)
	router.GET("/statistics/users/:userid", controllers.StatisticIndexUser)

	//Challenge
	router.POST("/challenges", controllers.ChallengeCreate)
	router.PUT("/challenges/:id", controllers.ChallengeUpdate)
	router.GET("/challenges", controllers.ChallengeIndex)
	router.GET("/challenges/:id", controllers.ChallengeShow)
	router.DELETE("/challenges/:id", controllers.ChallengeDelete)
	router.GET("/challenges/users/:userid", controllers.ChallengeShow)

	//Submission
	router.POST("/submissions", controllers.SubmissionCreate)
	router.PUT("/submissions/:id", controllers.SubmissionUpdate)
	router.GET("/submissions", controllers.SubmissionIndex)
	router.GET("/submissions/:id", controllers.SubmissionShow)
	router.DELETE("/submissions/:id", controllers.SubmissionDelete)
	router.GET("/submissions/id", controllers.SubmissionShowIndex)

	router.Run()
}
