package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mobintmu/golang_crud/internal/models"
	"github.com/mobintmu/golang_crud/internal/repositories"
	"github.com/mobintmu/golang_crud/internal/services"
)

// UserRoutes loads user-related endpoints
func UserRoutes(rg *gin.RouterGroup) {
	userService := services.NewUserService(repositories.UserRepository{})

	rg.GET("/", func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	rg.POST("", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := userService.CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Status(201)
	})
}
