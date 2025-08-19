package main

import (
	"gin/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Register(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
