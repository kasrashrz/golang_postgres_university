package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/Controllers"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func main() {
	server := gin.Default()
	db := models.SetupModels() // new
	server.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})
	server.GET("/", controllers.FindStudent)
	server.POST("/stu", func(ctx *gin.Context) {
		var student models.CreateStudentInput
		ctx.BindJSON(&student)
		fmt.Print(student)
	})

	server.Run()
}
