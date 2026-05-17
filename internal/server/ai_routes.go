package server

import (
	"github.com/gin-gonic/gin"

	h "tbd-backend/internal/handlers"
)

func MapAIRoutes(router *gin.RouterGroup, handlers *h.AIHandler) {
	ai_routes := router.Group("/ai")

	ai_routes.GET("/chat", handlers.HandleChat)
}
