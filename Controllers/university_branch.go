package Controllers

import (
	"github.com/gin-gonic/gin"
	"iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
	"net/http"
	"gorm.io/gorm"
)


func FindUniversityBranch(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var uni_branches []models.UniversityBranch
	db.Find(&uni_branches)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : uni_branches,
	})
}

func CreateUniversityBranch(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input models.UniversityBranch
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	 UniBranch := models.UniversityBranch{
		 Name:         input.Name,
		 Address:      input.Address,
		 URL:          input.URL,
		 CreationDate: input.CreationDate,
		 UniversityID: input.UniversityID,
	 }
	for _, course := range input.Courses {
		UniBranch.Courses = append(UniBranch.Courses, course)
	}

	db.Save(&UniBranch)
	ctx.JSON(http.StatusOK,gin.H{
		"data" : true,
	})
}

func UpdateUniversityBranch(ctx *gin.Context) {
	var UniBranch models.UniversityBranch
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&UniBranch).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Teacher
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&UniBranch).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": UniBranch})
}

func DeleteUniversityBranch(ctx *gin.Context){
	db := ctx.MustGet("db").(gorm.DB)
	var uni_branch models.UniversityBranch
	if err := db.Where("id = ?", ctx.Param("id")).First(&uni_branch).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&uni_branch)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}