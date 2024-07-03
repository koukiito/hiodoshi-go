package main

import (
	"github.com/gin-gonic/gin"
	"hiodoshi-go/internal/handlers"
	"hiodoshi-go/internal/services"
)

func main() {
	//cfg := config.Load()

	r := gin.Default()

	hiodoshiService := services.NewHiodoshiService()
	hiodoshiHandler := handlers.NewHiodoshiHandler(hiodoshiService)

	api := r.Group("/api")
	{
		api.GET("/hiodoshi-go", hiodoshiHandler.GetHiodoshi)
	}

	//log.Fatal(r.Run(":" + cfg.Port))

	r.Run(":8080")
}