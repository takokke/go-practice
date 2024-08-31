package main

import (
	"myapp/controllers"
	"myapp/initializers"
	"myapp/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/sign_up", controllers.SignUp)
	router.POST("/sign_in", controllers.SignIn)
	//ログインしているかチェック
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run(":8000") // 0.0.0.0:8000 でサーバーを立てます。
}
