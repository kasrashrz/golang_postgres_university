package Controllers

import (
	"github.com/gin-gonic/gin"
	"iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

var branch models.UniversityBranch
func FindUniversityBranch(ctx *gin.Context){
	branch.READ(ctx)
}

func CreateUniversityBranch(ctx *gin.Context){
	branch.CREATE(ctx)
}

func UpdateUniversityBranch(ctx *gin.Context) {
	branch.UPDATE(ctx)
}

func DeleteUniversityBranch(ctx *gin.Context){
	branch.DELETE(ctx)
}

func FindStudentsByUniversityBranch(ctx *gin.Context) {
	branch.StudentByUniBranch(ctx)
}