package Controllers

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

var course models.Course

func FindCourse(ctx *gin.Context) {
	course.READ(ctx)
}

func CreateCourse(ctx *gin.Context) {
	course.CREATE(ctx)
}

func UpdateCourse(ctx *gin.Context) {
	course.UPDATE(ctx)
}

func DeleteCourse(ctx *gin.Context){
	course.DELETE(ctx)
}

func FindCourseByNameOrQuantiyPlace(ctx *gin.Context) {

}

/*
select courses.*
from courses
INNER JOIN university_branch_courses ON university_branch_courses.course_id=courses.id
INNER JOIN university_branches ON university_branches.id=university_branch_courses.university_branch_id
INNER JOIN universities ON universities.id = university_branches.university_id WHERE universities.id=1;
*/

func FindCourseByUni(ctx *gin.Context) {
	var course []models.Course
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN university_branch_courses ON university_branch_courses.course_id=courses.id").Joins("INNER JOIN university_branches ON university_branches.id=university_branch_courses.university_branch_id").Joins("INNER JOIN universities ON universities.id = university_branches.university_id").Where("universities.id = ?", ctx.Param("id")).Scan(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": course})
}
