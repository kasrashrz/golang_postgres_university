package Controllers

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
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
	course.CourseByNameOrQuantity(ctx)
}

/*
select courses.*
from courses
INNER JOIN university_branch_courses ON university_branch_courses.course_id=courses.id
INNER JOIN university_branches ON university_branches.id=university_branch_courses.university_branch_id
INNER JOIN universities ON universities.id = university_branches.university_id WHERE universities.id=1;
*/

func FindCourseByUni(ctx *gin.Context) {
	course.CourseByUniversity(ctx)
}
