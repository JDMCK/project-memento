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
	router.GET("/create", handlers.Create)

	// Partial routes
	router.POST("/api/search", func(c *gin.Context) {
		value := c.PostForm("search")
		c.String(http.StatusOK, "<h1>This is a test data thing! %s</h1>", value)
	})

	router.POST("/api/preview", func(c *gin.Context) {
		title := c.PostForm("post-title")
		body := c.PostForm("post-body")

		c.String(http.StatusOK, "<h2>%s</h2><p>%s</p>", title, body)
	})
}
