package routes

import (
	"bootstrap/internal/adapter/input/api/controller"
	"bootstrap/internal/adapter/output/note_http"
	"bootstrap/internal/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	noteClient := note_http.NewNoteClient()

	noteService := service.NewNoteService(noteClient)
	var noteController = controller.NewNoteController(noteService)

	apiGroup := r.Group("/api")
	apiGroup.GET("/notes", noteController.ListNotes)
}
