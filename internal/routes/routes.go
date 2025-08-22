package routes

import (
	"gin/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// Load templates folder
	router.LoadHTMLGlob("internal/templates/**/*")

	// Serve static files
	router.Static("/static", "./internal/static")

	// Public routes
	router.GET("/", handlers.Home)
	router.GET("/search", handlers.Search)

	// Partial routes
	router.POST("/api/search", func(c *gin.Context) {
		value := c.PostForm("search")
		c.String(http.StatusOK, "<h1>This is a test data thing! %s</h1>", value)
	})
}
