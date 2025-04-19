package routes

import (
	"bootstrap/internal/adapter/input/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	noteController := controller.NewNoteController()

	r.GET("/notes", noteController.ListNotes)
}
