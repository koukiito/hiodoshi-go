package main

import (
	"hiodoshi-go/internal/handlers"
	"hiodoshi-go/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	hiodoshiService := services.NewHiodoshiService()
	hiodoshiHandler := handlers.NewHiodoshiHandler(hiodoshiService)

	api := r.Group("/api")
	{
		api.GET("/hiodoshi-go", hiodoshiHandler.GetHiodoshi)
	}

	r.Run(":8080")
}
