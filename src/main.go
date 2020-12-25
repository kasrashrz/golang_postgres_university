package main

import (
	"github.com/gin-gonic/gin"
	controllers "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/Controllers"
	"iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func main() {
	server := gin.Default()
	db := models.SetupModels() // new
	server.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})
	server.GET("/", controllers.FindStudent)
	server.POST("/student", controllers.CreateStudent)
	server.DELETE("/student/delete/:id", controllers.DeleteStudent)

	server.Run()
}
