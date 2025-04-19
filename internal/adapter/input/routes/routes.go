package routes

import (
	"bootstrap/internal/adapter/input/controller"
	"bootstrap/internal/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	noteService := service.NewNoteService()
	var noteController = controller.NewNoteController(noteService)
	r.GET("/notes", noteController.ListNotes)
}
