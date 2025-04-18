package routes

import (
	"github.com/gin-gonic/gin"
)

// HealthRoutes loads health-related endpoints
func HealthRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
