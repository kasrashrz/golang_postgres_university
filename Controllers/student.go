package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	models "iran.gitlab.medrick.games/medrick/server/go_lang/sample_postgres_project/models"
)

func FindStudent(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var students []models.Students
	db.Find(&students)
	ctx.JSON(http.StatusOK, gin.H{"result": students})
}
