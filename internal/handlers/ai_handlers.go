package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	aiServices int
}

func NewAIHandler(service int) *AIHandler {
	return &AIHandler{aiServices: 0}
}

// This is just here for base setup testing
func (h *AIHandler) HandleChat(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
