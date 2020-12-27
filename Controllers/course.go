package Controllers

import (
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

//func CreateCourse(ctx *gin.Context){
//	db := ctx.MustGet("db").(*gorm.DB)
//	var input models.CreateCourseInput
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	newCourse := models.CreateCourseInput{}
//	db.Save(&newCourse)
//	ctx.JSON(http.StatusOK,gin.H{
//		"data" : true,
//	})
//}

//func UpdateTeacher(ctx *gin.Context) {
//	var teacher models.UpdateTeacherInput
//	db := ctx.MustGet("db").(*gorm.DB)
//	if err := db.Where("id = ?", ctx.Param("id")).First(&teacher).Error; err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
//		return
//	}
//	var input models.UpdateTeacherInput
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	db.Model(&teacher).Updates(input)
//	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
//}

func DeleteCourse(ctx *gin.Context){
	db := ctx.MustGet("db").(gorm.DB)
	var course models.Course
	if err := db.Where("id = ?", ctx.Param("id")).First(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&course)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}