package main

import (
	"github.com/NCInside/procod_01/controllers"
	"github.com/NCInside/procod_01/initializers"
	"github.com/NCInside/procod_01/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	//User
	users := router.Group("/users")
	{
		users.POST("/", controllers.UserCreate)
		users.PUT("/:id", controllers.UserUpdate)
		users.GET("/", controllers.UserIndex)
		users.GET("/:id", controllers.UserShow)
		users.DELETE("/:id", controllers.UserDelete)
	}

	//Statistic
	statistics := router.Group("/statistics")
	{
		statistics.POST("/", controllers.StatisticCreate)
		statistics.PUT("/:id", controllers.StatisticUpdate)
		statistics.GET("/", controllers.StatisticIndex)
		statistics.GET("/:id", controllers.StatisticShow)
		statistics.DELETE("/:id", controllers.StatisticDelete)
		statistics.GET("/users/:userid", controllers.StatisticIndexUser)
	}

	//Challenge
	challenges := router.Group("/challenges")
	{
		challenges.POST("/", controllers.ChallengeCreate)
		challenges.PUT("/:id", controllers.ChallengeUpdate)
		challenges.GET("/", controllers.ChallengeIndex)
		challenges.GET("/:id", controllers.ChallengeShow)
		challenges.DELETE("/:id", controllers.ChallengeDelete)
		challenges.GET("/users/:userid", controllers.ChallengeIndexUser)

		//ChallengeLabel
		label := challenges.Group("/label")
		{
			label.POST("/", controllers.ChallengeLbCreate)
			label.PUT("/:id", controllers.ChallengeLbUpdate)
			label.GET("/", controllers.ChallengeLbIndex)
			label.GET("/:id", controllers.ChallengeLbShow)
			label.DELETE("/:id", controllers.ChallengeLbDelete)
			label.GET("/challenge/:challengeid", controllers.ChallengeLbIndexChallenge)
		}

		//ChallengeTarget
		target := challenges.Group("/target")
		{
			target.POST("/", controllers.ChallengeTrCreate)
			target.PUT("/:id", controllers.ChallengeTrUpdate)
			target.GET("/", controllers.ChallengeTrIndex)
			target.GET("/:id", controllers.ChallengeTrShow)
			target.DELETE("/:id", controllers.ChallengeTrDelete)
			target.GET("/challenge/:challengeid", controllers.ChallengeTrIndexChallenge)
		}

		//ChallengeExample
		example := challenges.Group("/example")
		{
			example.POST("/", controllers.ChallengeExCreate)
			example.PUT("/:id", controllers.ChallengeExUpdate)
			example.GET("/", controllers.ChallengeExIndex)
			example.GET("/:id", controllers.ChallengeExShow)
			example.DELETE("/:id", controllers.ChallengeExDelete)
			example.GET("/challenge/:challengeid", controllers.ChallengeExIndexChallenge)

		}
	}

	//Submission
	submissions := router.Group("/submissions")
	{
		submissions.POST("/", controllers.SubmissionCreate)
		submissions.PUT("/:id", controllers.SubmissionUpdate)
		submissions.GET("/", controllers.SubmissionIndex)
		submissions.GET("/:id", controllers.SubmissionShow)
		submissions.DELETE("/:id", controllers.SubmissionDelete)
		submissions.GET("/id", controllers.SubmissionShowIndex)
	}

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	router.Run()
}
