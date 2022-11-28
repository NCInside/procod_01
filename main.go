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

	router.POST("/users", controllers.UserCreate)
	router.PUT("/users/:id", controllers.UserUpdate)
	router.GET("/users", controllers.UserIndex)
	router.GET("/users/:id", controllers.UserShow)
	router.DELETE("/users/:id", controllers.UserDelete)

	router.POST("/statistics", controllers.StatisticCreate)
	router.PUT("/statistics/:id", controllers.StatisticUpdate)
	router.GET("/statistics", controllers.StatisticIndex)
	router.GET("/statistics/:id", controllers.StatisticShow)
	router.DELETE("/statistics/:id", controllers.StatisticDelete)
	router.GET("/statistics/:userid", controllers.StatisticIndexUser)

	router.Run()
}
