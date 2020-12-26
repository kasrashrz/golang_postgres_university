package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	"github.com/jinzhu/gorm"
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
	var input models.CreateStudentInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStudent := models.Student{
		Model:        gorm.Model{},
		Name:         input.Name,
		Age:          input.Age,
		Mail:         input.Mail,
		NationalCode: input.NationalCode,
		Address:      input.Address,

		//Courses: []models.Course{{Model:gorm.Model{} , Name: input.Courses, QuantityPlace: , StartDate: , EndDate: , CreatedDate: , Students: }},
	}
	db.Create(&newStudent)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})

}
func UpdateStudent(ctx *gin.Context) {
	var student models.Student
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&student).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
