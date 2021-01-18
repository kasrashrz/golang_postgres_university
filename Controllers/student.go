package Controllers

import (
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

var student models.Student

func FindStudent(ctx *gin.Context) {
	student.READ(ctx)
}

func CreateStudent(ctx *gin.Context) {
	student.CREATE(ctx)
}

func UpdateStudent(ctx *gin.Context) {
	student.UPDATE(ctx)
}

func DeleteStudent(ctx *gin.Context) {
	student.DELETE(ctx)
}

func FindStudentByNameOrMail(ctx *gin.Context) {
	student.StudentByNameOrMail(ctx)
}