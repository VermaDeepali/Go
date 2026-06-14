package handler

import (
	"net/http"

	"ai-gateway-system-go/internal/model"
	"ai-gateway-system-go/internal/service"

	"github.com/gin-gonic/gin"
)

func AskHandler(c *gin.Context) {
	var req model.AskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	response, err := service.ProcessRequest(req.Prompt, req.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
