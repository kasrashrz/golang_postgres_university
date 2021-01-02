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

func CreateUniversity(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input models.CreateUniversityInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUniversity := models.CreateUniversityInput{
		Name:             input.Name,
		Address:          input.Address,
		URL:              input.URL,
		CreationDate:     input.CreationDate,
	}
	for _, student := range input.Students {
		newUniversity.Students = append(newUniversity.Students, student)
	}
	for _, uni_branches := range input.UniversityBranches {
		newUniversity.UniversityBranches = append(newUniversity.UniversityBranches, uni_branches)
	}
	db.Save(&newUniversity)
	ctx.JSON(http.StatusOK,gin.H{
		"data" : true,
	})
}

func UpdateUniversity(ctx *gin.Context) {
	var university models.UpdateUniversityInput
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&university).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdateUniversityInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&university).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": university})
}

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