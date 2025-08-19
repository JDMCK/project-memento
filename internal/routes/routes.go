package routes

import (
	"gin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// Load templates folder
	router.LoadHTMLGlob("internal/templates/**/*")

	// Public routes
	router.GET("/", handlers.Home)
	router.GET("/blog", handlers.Blog)
}
