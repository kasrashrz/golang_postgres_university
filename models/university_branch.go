package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UniversityBranch struct {
	gorm.Model
	Name         string   `gorm:"unique;"json:"name"`
	Address      string   `json:"address"`
	URL 		 string	  `json:"url"`
	CreationDate string	  `json:"CreationDate"`
	UniversityID uint
	University University `gorm:"references:ID"`
	Courses      []Course `gorm:"many2many:UniversityBranch_courses"`
}

type Uni_Branch_Handler interface {
	CREATE(ctx *gin.Context)
	READ(ctx *gin.Context)
	UPDATE(ctx *gin.Context)
	DELETE(ctx *gin.Context)
	StudentByNameOrMail(ctx *gin.Context)
}

func (uni_branch *UniversityBranch) CREATE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input UniversityBranch
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	UniBranch := UniversityBranch{
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
	fmt.Println("CREATE INTERFACE")
}
func (uni_branch *UniversityBranch) READ(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var uni_branches []UniversityBranch
	db.Find(&uni_branches)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : uni_branches,
	})
	fmt.Println("READ INTERFACE")
}
func (uni_branch UniversityBranch) UPDATE(ctx *gin.Context){
	var UniBranch UniversityBranch
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&UniBranch).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input Teacher
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&UniBranch).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": UniBranch})
	fmt.Println("UPDATE INTERFACE")
}
func (branch *UniversityBranch) DELETE(ctx *gin.Context){
	db := ctx.MustGet("db").(gorm.DB)
	var uni_branch UniversityBranch
	if err := db.Where("id = ?", ctx.Param("id")).First(&uni_branch).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&uni_branch)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})

	fmt.Println("DELETE INTERFACE")
}
func (uni_branch *UniversityBranch) StudentByUniBranch(ctx *gin.Context){
	//var university models.University
	var student []Student

	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("students").Select("*").Where("university_branch_id = ?", ctx.Param("id")).Scan(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": student})
	fmt.Println("STU API")
}