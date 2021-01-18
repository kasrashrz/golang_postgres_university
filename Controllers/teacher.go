package Controllers

import (
	"github.com/gin-gonic/gin"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

var teacher models.Teacher
func FindTeacher(ctx *gin.Context){
	teacher.READ(ctx)
}

func CreateTeacher(ctx *gin.Context){
	teacher.CREATE(ctx)
}

func UpdateTeacher(ctx *gin.Context) {
	teacher.UPDATE(ctx)
}

func DeleteTeacher(ctx *gin.Context) {
	teacher.DELETE(ctx)
}

func FindTeacherStudents(ctx *gin.Context) {
	teacher.Student_teacher(ctx)
}