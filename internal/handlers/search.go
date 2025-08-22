package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", nil)
}
