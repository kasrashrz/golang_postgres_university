package Controllers

import (
	"github.com/gin-gonic/gin"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

var university models.University

func FindUniversity(ctx *gin.Context){
	university.READ(ctx)
}

func CreateUniversity(ctx *gin.Context){
	university.CREATE(ctx)
}

func UpdateUniversity(ctx *gin.Context) {
	university.UPDATE(ctx)
}

func DeleteUniversity(ctx *gin.Context){
	university.DELETE(ctx)
}

/*
SELECT students.*, university_branches.*, universities.*
FROM ((students
INNER JOIN university_branches ON students.university_branch_id = university_branches.id)
INNER JOIN universities ON university_branches.university_id = universities.id and universities.id=REQUESTED UNI ID);
*/

func FindStudentsByUniversity(ctx *gin.Context) {
	university.StudentByUniversity(ctx)
}