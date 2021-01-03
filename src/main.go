package main

import (
	"github.com/gin-gonic/gin"
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
	//server.GET("/students", controllers.FindStudent)
	//server.POST("/student/create", controllers.CreateStudent)
	//server.POST("/student/delete/:id", controllers.DeleteStudent)
	//server.POST("/student/update/:id", controllers.UpdateStudent)
	//// teachers
	//server.GET("/teachers", controllers.FindTeacher)
	//server.POST("/teacher/create", controllers.CreateTeacher)
	//server.POST("/teacher/delete/:id", controllers.DeleteTeacher)
	//server.POST("/teacher/update/:id", controllers.UpdateTeacher)
	//// courses
	//server.GET("/courses", controllers.FindCourse)
	//server.POST("/course/create", controllers.CreateCourse)
	//server.POST("/course/delete/:id", controllers.DeleteCourse)
	//server.POST("/course/update/:id", controllers.UpdateCourse)
	//// universities
	//server.GET("/universities", controllers.FindUniversity)
	//server.POST("/university/create", controllers.CreateUniversity)
	//server.POST("/university/delete/:id", controllers.DeleteUniversity)
	//server.POST("/university/update/:id", controllers.UpdateUniversity)
	//// university_branches
	//server.GET("/university_branches", controllers.FindUniversityBranch)
	//server.POST("/university_branches/create", controllers.CreateUniversityBranch)
	//server.POST("/university_branches/delete/:id", controllers.DeleteUniversityBranch)
	//server.POST("/university_branches/update/:id", controllers.UpdateUniversityBranch)
	//


	server.Run()
}
