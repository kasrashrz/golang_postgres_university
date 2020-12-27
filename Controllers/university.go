package Controllers

import (
"github.com/gin-gonic/gin"
"gorm.io/gorm"
models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
"net/http"
)

func FindUniversity(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var universities []models.University
	db.Find(&universities)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : universities,
	})
}

//func CreateTeacher(ctx *gin.Context){
//	db := ctx.MustGet("db").(*gorm.DB)
//	var input models.CreateTeacherInput
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	newTeacher := models.CreateTeacherInput{}
//	db.Save(&newTeacher)
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

func DeleteUniversity(ctx *gin.Context){
	db := ctx.MustGet("db").(gorm.DB)
	var university models.University
	if err := db.Where("id = ?", ctx.Param("id")).First(&university).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&university)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}