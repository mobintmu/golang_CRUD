package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the main router
func SetupRouter() *gin.Engine {
	r := gin.New()

	// Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Mount routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Load route groups
	LoadRouteGroups(r)

	return r
}

// LoadRouteGroups mounts all route groups
func LoadRouteGroups(r *gin.Engine) {
	// Load health routes
	HealthRoutes(r.Group("/api"))

	UserRoutes(r.Group("/api/users"))

}
