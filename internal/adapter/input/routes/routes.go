package routes

import (
	"bootstrap/internal/adapter/input/controller"
	"bootstrap/internal/adapter/output/note_http"
	"bootstrap/internal/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	noteClient := note_http.NewNoteClient()

	noteService := service.NewNoteService(noteClient)
	var noteController = controller.NewNoteController(noteService)
	r.GET("/notes", noteController.ListNotes)
}
