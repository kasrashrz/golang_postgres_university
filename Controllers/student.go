package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func FindStudent(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var students []models.Student
	db.Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"result": students})
}

func CreateStudent(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	//fmt.Print("state1")
	var input models.Student
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//fmt.Print("state2")
		return
	}
	fmt.Print(input)
	newStudent := models.Student{
		Name:         input.Name,
		Age:          input.Age,
		Mail:         input.Mail,
		NationalCode: input.NationalCode,
		Address:      input.Address,
		UniversityID: input.UniversityID,
	}
	for _, course := range input.Courses {
		newStudent.Courses = append(newStudent.Courses, course)
	}
	db.Save(&newStudent)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})

}

func UpdateStudent(ctx *gin.Context) {
	var student models.Student
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Student
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, course := range input.Courses {
		student.Courses = append(student.Courses, course)
	}
	db.Model(&student).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": student})
}

func DeleteStudent(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var student models.Student
	if err := db.Where("id = ?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db.Delete(&student)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})
}
