package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
	"net/http"
)

func FindTeacher(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var teachers []models.Teacher
	db.Find(&teachers)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : teachers,
	})
}

func CreateTeacher(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input models.CreateTeacherInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTeacher := models.CreateTeacherInput{
		Name:         input.Name,
		Mail:         input.Mail,
		NationalCode: input.NationalCode,
	}

	for _, student := range input.Students {
		newTeacher.Students = append(newTeacher.Students, student)
	}
	for _, course := range input.Courses {
		newTeacher.Courses = append(newTeacher.Courses, course)
	}

	db.Save(&newTeacher)
	ctx.JSON(http.StatusOK,gin.H{
		"data" : true,
	})
}

func UpdateTeacher(ctx *gin.Context) {
	var teacher models.Teacher
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&teacher).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Teacher
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&teacher).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
}

func DeleteTeacher(ctx *gin.Context){
	db := ctx.MustGet("db").(gorm.DB)
	var teacher models.Teacher
	if err := db.Where("id = ?", ctx.Param("id")).First(&teacher).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&teacher)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}