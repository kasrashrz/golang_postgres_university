package models

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name               string `gorm:"unique;" json:"name"`
	QuantityPlace      int    `json:"quantityPlace"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"EndDate"`
	CreatedDate        string `json:"CreatedDate"`
	TeacherID          uint
	Teacher            Teacher            `gorm:"references:ID"`
	UniversityBranches []UniversityBranch `gorm:"many2many:UniversityBranch_courses;"`
	Students           []Student          `gorm:"many2many:students_courses;"`
}


type Course_Handler interface {
	CONNECT(ctx *gin.Context)
	CREATE(ctx *gin.Context)
	READ(ctx *gin.Context)
	UPDATE(ctx *gin.Context)
	DELETE(ctx *gin.Context)
	CourseByNameOrQuantity(ctx *gin.Context)
	CourseByUniversity(ctx *gin.Context)
}

func (course *Course) CREATE(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var input Course
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(input)
	newCourse := Course{
		Name:          input.Name,
		QuantityPlace: input.QuantityPlace,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		CreatedDate:   input.CreatedDate,
		TeacherID:     input.TeacherID,
	}
	for _, student := range input.Students {
		newCourse.Students = append(newCourse.Students, student)
	}
	for _, university_branches := range input.UniversityBranches {
		newCourse.UniversityBranches = append(newCourse.UniversityBranches, university_branches)
	}
	db.Save(&newCourse)
	ctx.JSON(http.StatusAccepted, gin.H{"data": true})
	fmt.Println("CREATE INTERFACE")
}
func (input *Course) READ(ctx *gin.Context)   {
	db := ctx.MustGet("db").(*gorm.DB)
	var courses []Course
	db.Find(&courses)
	ctx.JSON(http.StatusOK, gin.H{
		"result": courses,
	})
	fmt.Println("READ INTERFACE")

}
func (course Course) UPDATE(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ctx.Param("id")).First(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input Course
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&course).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": course})
	fmt.Println("UPDATE INTERFACE")
}
func (input *Course) DELETE(ctx *gin.Context) {
		db := ctx.MustGet("db").(*gorm.DB)
		var course Course
		if err := db.Where("id = ?", ctx.Param("id")).First(&course).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		db.Delete(&course)
		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
		fmt.Println("DELETE INTERFACE")
}
func (course Course) CourseByNameOrQuantity(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("courses").Select("*").Where("quantity_place = ?", ctx.Param("name")).Scan(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": course})
	fmt.Println("COURSE_NAME_OR_QUANTITY API")

}
func (course Course) CourseByUniversity(ctx *gin.Context){
	db := ctx.MustGet("db").(*gorm.DB)
	if err := db.Table("courses").Select("courses.*").Joins("INNER JOIN university_branch_courses ON university_branch_courses.course_id=courses.id").Joins("INNER JOIN university_branches ON university_branches.id=university_branch_courses.university_branch_id").Joins("INNER JOIN universities ON universities.id = university_branches.university_id").Where("universities.id = ?", ctx.Param("id")).Scan(&course).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": course})
	fmt.Println("COURSE_UNI API")
}