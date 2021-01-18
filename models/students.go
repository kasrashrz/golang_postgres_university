package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Student struct {
	gorm.Model
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique"json:"nationalCode"`
	Address      string   `json:"address"`
	Courses      []Course `gorm:"many2many:students_courses;"`
	//Teachers      []Teacher `gorm:"many2many:students_teachers;"`
	UniversityBranchID	uint 	`json:"university_id"`
	UniversityBranch UniversityBranch `gorm:"references:ID"`
}


type Student_Handler interface {
	CREATE(ctx *gin.Context)
	READ(ctx *gin.Context)
	UPDATE(ctx *gin.Context)
	DELETE(ctx *gin.Context)
	StudentByNameOrMail(ctx *gin.Context)
}

func (student *Student) CREATE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	//fmt.Print("state1")
	var input Student
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//fmt.Print("state2")
		return
	}
	fmt.Print(input)
	newStudent := Student{
		Name:         input.Name,
		Age:          input.Age,
		Mail:         input.Mail,
		NationalCode: input.NationalCode,
		Address:      input.Address,
		UniversityBranchID: input.UniversityBranchID,
	}
	for _, course := range input.Courses {
		newStudent.Courses = append(newStudent.Courses, course)
	}
	db.Save(&newStudent)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})
	fmt.Println("CREATE INTERFACE")
}
func (student *Student) READ(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var students []Student
	db.Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"result": students})
	fmt.Println("READ INTERFACE")
}
func (student *Student) UPDATE(ctx *gin.Context){
	//var student models.Student
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input Student
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, course := range input.Courses {
		student.Courses = append(student.Courses, course)
	}
	db.Model(&student).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": student})
	fmt.Println("UPDATE INTERFACE")
}
func (student *Student) DELETE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&student)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})
	fmt.Println("DELETE INTERFACE")
}
func (input *Student) StudentByNameOrMail(ctx *gin.Context){
	var student []Student
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("students").Select("*").Where("name =? OR mail = ?", ctx.Param("name"), ctx.Param("name")).Scan(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": student})
	fmt.Println("STU API")
}