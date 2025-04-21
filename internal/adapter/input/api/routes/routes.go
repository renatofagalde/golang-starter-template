package routes

import (
	"bootstrap/internal/adapter/input/api/controller"
	"bootstrap/internal/adapter/output/factory"
	"bootstrap/internal/adapter/output/note_http"
	"bootstrap/internal/application/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.RouterGroup, database *gorm.DB) {

	httpClient := note_http.NewNoteClient()

	noteFactory := factory.NewNoteFactory(database, httpClient)
	noteService := service.NewNoteService(noteFactory)

	var noteController = controller.NewNoteController(noteService)

	apiGroup := r.Group("/api")
	apiGroup.GET("/notes", noteController.ListNotes)
}
