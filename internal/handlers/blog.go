package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Blog(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", nil)
}
