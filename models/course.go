package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	CONNECT()
	CREATE(ctx *gin.Context)
	READ()
	UPDATE()
	DELETE()
}

func (course *Course) CREATE(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var courses []Course
	db.Find(&courses)
	ctx.JSON(http.StatusOK, gin.H{
		"result": courses,
	})
	fmt.Print("Create Func")
}
func (input *Course) READ()   {}
func (input *Course) UPDATE() {}
func (input *Course) DELETE() {}
