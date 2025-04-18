package services

import (
	"github.com/gin-gonic/gin"
	"github.com/mobintmu/golang_crud/internal/models"
	"github.com/mobintmu/golang_crud/internal/repositories"
)

// UserService handles user-related operations
type UserService struct {
	repository repositories.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repository.GetAllUsers()
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repository.GetUser(id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.repository.CreateUser(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id string, user *models.User) error {
	return s.repository.UpdateUser(id, user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id string) error {
	return s.repository.DeleteUser(id)
}

// UserRoutes returns router group for user endpoints
func UserRoutes(rg *gin.RouterGroup, userService *UserService) {
	rg.GET("/", func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	rg.POST("", func(c *gin.Context) {
		user, err := models.BindUser(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = userService.CreateUser(user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Status(201)
	})
}
