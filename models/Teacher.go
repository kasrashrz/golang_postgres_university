package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Teacher struct {
	gorm.Model
	Name          string  `json:"name"`
	Mail         string   `json:"mail"`
	NationalCode string   `gorm:"unique;" json:"nationalCode"`
}


type Teacher_Handler interface {
	CREATE(ctx *gin.Context)
	READ(ctx *gin.Context)
	UPDATE(ctx *gin.Context)
	DELETE(ctx *gin.Context)
	Student_teacher(ctx *gin.Context)
}

func (teacher *Teacher) CREATE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var input Teacher
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTeacher := Teacher{
		Name:         input.Name,
		Mail:         input.Mail,
		NationalCode: input.NationalCode,
	}
	db.Save(&newTeacher)
	ctx.JSON(http.StatusOK,gin.H{
		"data" : true,
	})
	fmt.Println("CREATE INTERFACE")
}
func (teacher *Teacher) READ(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	var teachers []Teacher
	db.Find(&teachers)
	ctx.JSON(http.StatusOK,gin.H{
		"result" : teachers,
	})
	fmt.Println("READ INTERFACE")

}
func (teacher Teacher) UPDATE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&teacher).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input Teacher
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&teacher).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": teacher})
	fmt.Println("UPDATE INTERFACE")

}
func (teacher *Teacher) DELETE(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&teacher).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&teacher)
	ctx.JSON(http.StatusOK,gin.H{
		"data":true,
	})
	fmt.Println("DELETE INTERFACE")

}
func (teacher *Teacher) Student_teacher(ctx *gin.Context){
	/*
	   SELECT students.*
	   FROM students
	   INNER JOIN students_courses ON students.id = students_courses.student_id
	   INNER JOIN courses ON courses.id = students_courses.course_id INNER JOIN teachers ON teachers.id=courses.teacher_id where teachers.id=1;
	*/
	var student []Student
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("students").Select("students.*").Joins("INNER JOIN students_courses ON students.id = students_courses.student_id").Joins("INNER JOIN courses ON courses.id = students_courses.course_id").Joins("INNER JOIN teachers ON teachers.id=courses.teacher_id").Where("teachers.id = ?", ctx.Param("id")).Scan(&student).Error; err != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": student})
	fmt.Println("API")
}