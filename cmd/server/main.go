package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hftamayo/gorptmservice/internal/handlers"
)

func main() {
	router := gin.Default()

	router.POST("/generate-report", handlers.GenerateReportHandler)

	router.Run(":8004") // Listen and serve on 0.0.0.0:8080
}
