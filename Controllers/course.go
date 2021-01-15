package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
	"net/http"
)

func FindCourse(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var courses []models.Course
	db.Find(&courses)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : courses,
	})
}

func CreateCourse(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input models.Course
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(input)
	newCourse := models.Course{
		Name:             input.Name,
		QuantityPlace:    input.QuantityPlace,
		StartDate:        input.StartDate,
		EndDate:          input.EndDate,
		CreatedDate:      input.CreatedDate,
		TeacherID: input.TeacherID,
	}
	for _, student := range input.Students{
		newCourse.Students = append(newCourse.Students, student)
	}
	for _, university_branches := range input.UniversityBranches{
		newCourse.UniversityBranches = append(newCourse.UniversityBranches, university_branches)
	}
	db.Save(&newCourse)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})
}

func UpdateCourse(ctx *gin.Context) {
	var course models.Course
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Course
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&course).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": course})
}

func DeleteCourse(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var course models.Course
	if err := db.Where("id = ?", ctx.Param("id")).First(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&course)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}

func FindCourseByNameOrQuantiyPlace(ctx *gin.Context){
	var course []models.Course
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("courses").Select("*").Where("quantity_place = ?", ctx.Param("name")).Scan(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": course})
}

/*
select courses.*
from courses
INNER JOIN university_branch_courses ON university_branch_courses.course_id=courses.id
INNER JOIN university_branches ON university_branches.id=university_branch_courses.university_branch_id
INNER JOIN universities ON universities.id = university_branches.university_id WHERE universities.id=1;
*/

func FindCourseByUni(ctx *gin.Context){
	var course []models.Course
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN university_branch_courses ON university_branch_courses.course_id=courses.id").Joins("INNER JOIN university_branches ON university_branches.id=university_branch_courses.university_branch_id").Joins("INNER JOIN universities ON universities.id = university_branches.university_id").Where("universities.id = ?", ctx.Param("id")).Scan(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": course})
}