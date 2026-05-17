package server

import (
	"github.com/gin-gonic/gin"

	h "tbd-backend/internal/handlers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1ApiRouter := router.Group("/api/v1")

	aiHandler := h.NewAIHandler(0)
	MapAIRoutes(v1ApiRouter, aiHandler)

	return router
}
