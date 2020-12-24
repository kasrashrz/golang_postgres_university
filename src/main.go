package main

import (
	"github.com/gin-gonic/gin"
	"iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func main() {
	server := gin.Default()
	db := models.SetupModels() // new
	server.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	server.GET("/", controllers.findStudent())
	server.Run()
}
