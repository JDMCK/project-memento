package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}
