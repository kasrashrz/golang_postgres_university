package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func findStudent(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var students []models.Students
	db.Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"data": students})
}
