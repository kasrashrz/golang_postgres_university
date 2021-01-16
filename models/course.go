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
func (input *Course) UPDATE(ctx *gin.Context) {}
func (input *Course) DELETE(ctx *gin.Context) {}
