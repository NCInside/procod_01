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
	router.GET("/challenges/users/:userid", controllers.ChallengeIndexUser)

	//ChallengeLabel
	router.POST("/challenges/label", controllers.ChallengeLbCreate)
	router.PUT("/challenges/label/:id", controllers.ChallengeLbUpdate)
	router.GET("/challenges/label", controllers.ChallengeLbIndex)
	router.GET("/challenges/label/:id", controllers.ChallengeLbShow)
	router.DELETE("/challenges/label/:id", controllers.ChallengeLbDelete)
	router.GET("/challenges/label/challenge/:challengeid", controllers.ChallengeLbIndexChallenge)

	//ChallengeTarget
	router.POST("/challenges/target", controllers.ChallengeTrCreate)
	router.PUT("/challenges/target/:id", controllers.ChallengeTrUpdate)
	router.GET("/challenges/target", controllers.ChallengeTrIndex)
	router.GET("/challenges/target/:id", controllers.ChallengeTrShow)
	router.DELETE("/challenges/target/:id", controllers.ChallengeTrDelete)
	router.GET("/challenges/target/challenge/:challengeid", controllers.ChallengeTrIndexChallenge)

	//ChallengeExample
	router.POST("/challenges/example", controllers.ChallengeExCreate)
	router.PUT("/challenges/example/:id", controllers.ChallengeExUpdate)
	router.GET("/challenges/example", controllers.ChallengeExIndex)
	router.GET("/challenges/example/:id", controllers.ChallengeExShow)
	router.DELETE("/challenges/example/:id", controllers.ChallengeExDelete)
	router.GET("/challenges/example/challenge/:challengeid", controllers.ChallengeExIndexChallenge)

	//Submission
	router.POST("/submissions", controllers.SubmissionCreate)
	router.PUT("/submissions/:id", controllers.SubmissionUpdate)
	router.GET("/submissions", controllers.SubmissionIndex)
	router.GET("/submissions/:id", controllers.SubmissionShow)
	router.DELETE("/submissions/:id", controllers.SubmissionDelete)
	router.GET("/submissions/id", controllers.SubmissionShowIndex)

	router.Run()
}
