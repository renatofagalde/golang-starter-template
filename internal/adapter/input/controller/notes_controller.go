package controller

import "github.com/gin-gonic/gin"

type noteController struct{}

func NewNoteController() *noteController {
	return &noteController{}
}

func (c *noteController) ListNotes(ctx *gin.Context) {

}
