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
	var input models.University
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUniversity := models.University{
		Name:             input.Name,
		Address:          input.Address,
		URL:              input.URL,
		CreationDate:     input.CreationDate,

	}
	db.Save(&newUniversity)
	ctx.JSON(http.StatusOK,gin.H{
		"data" : true,
	})
}

func UpdateUniversity(ctx *gin.Context) {
	var university models.University
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&university).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.University
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&university).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": university})
}

func DeleteUniversity(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var university models.University
	if err := db.Where("id = ?", ctx.Param("id")).First(&university).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&university)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
}

/*
SELECT students.*, university_branches.*, universities.*
FROM ((students
INNER JOIN university_branches ON students.university_branch_id = university_branches.id)
INNER JOIN universities ON university_branches.university_id = universities.id and universities.id=REQUESTED UNI ID);
*/

func FindStudentsByUniversity(ctx *gin.Context) {
	var student []models.Student

	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("students").Select("*").Where("university_branch_id = ?", ctx.Param("id")).Scan(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": student})
}