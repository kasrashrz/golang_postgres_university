package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type University struct {
	gorm.Model
	Name         string   `json:"name"`
	Address      string   `json:"address"`
	URL 		 string	  `json:"url"`
    CreationDate string	  `json:"CreationDate"`
}

type University_Handler interface {
	CREATE(ctx *gin.Context)
	READ(ctx *gin.Context)
	UPDATE(ctx *gin.Context)
	DELETE(ctx *gin.Context)
	StudentByNameOrMail(ctx *gin.Context)
}

func (university *University) CREATE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input University
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUniversity := University{
		Name:             input.Name,
		Address:          input.Address,
		URL:              input.URL,
		CreationDate:     input.CreationDate,

	}
	db.Save(&newUniversity)
	ctx.JSON(http.StatusOK,gin.H{
		"data" : true,
	})
	fmt.Println("CREATE INTERFACE")
}
func (university *University) READ(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var universities []University
	db.Find(&universities)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : universities,
	})
	fmt.Println("READ INTERFACE")
}
func (university University) UPDATE(ctx *gin.Context) {
	//var university models.University
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&university).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input University
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&university).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": university})
	fmt.Println("UPDATE INTERFACE")
}
func (university University) DELETE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	//var university models.University
	if err := db.Where("id = ?", ctx.Param("id")).First(&university).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&university)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
	fmt.Println("DELETE INTERFACE")
}
func (university *University) StudentByUniversity(ctx *gin.Context){
	var student []Student
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("students").Select("students.*").Joins("INNER JOIN university_branches ON students.university_branch_id = university_branches.id").Joins("INNER JOIN universities ON university_branches.university_id = universities.id and universities.id=2").Scan(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": student})
	fmt.Println("STU API")
}