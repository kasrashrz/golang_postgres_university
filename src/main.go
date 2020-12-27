package main

import (
	"github.com/gin-gonic/gin"
	controllers "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/Controllers"
	"iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func main() {
	server := gin.Default()
	db := models.SetupModels() // Create models and set the first one manual
	server.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})
	// students
	server.GET("/students", controllers.FindStudent)
	server.POST("/student/create", controllers.CreateStudent)
	//server.POST("/student/delete/:id", controllers.DeleteStudent)
	server.POST("/student/update/:id", controllers.UpdateStudent)
	// teachers
	server.GET("/teachers", controllers.FindTeacher)
	//server.POST("/teacher/create", controllers.)
	//server.POST("/teacher/delete/:id", controllers.)
	server.POST("/teacher/delete/:id", controllers.DeleteTeacher)
	// courses
	server.GET("/courses", controllers.FindCourse)
	//server.POST("/teacher/create", controllers.)
	//server.POST("/teacher/delete/:id", controllers.)
	server.POST("/course/delete/:id", controllers.DeleteCourse)

	server.Run()
}
