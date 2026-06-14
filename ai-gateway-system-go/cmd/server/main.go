package main

import (
	"ai-gateway-system-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/v1/ai/ask", handler.AskHandler)

	r.Run(":8080")
}
