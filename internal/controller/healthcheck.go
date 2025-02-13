package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckResponse defines the response structure for the health check.
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HealthCheck(ctx *gin.Context) {
	response := HealthCheckResponse{
		Status:  "healthy",
		Message: "All systems operational",
	}
	ctx.JSON(http.StatusOK, response)
}
