package models

import (
	"github.com/gin-gonic/gin"
)

// User represents a user entity
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// BindUser binds JSON body to User struct
func BindUser(c *gin.Context) (*User, error) {
	var u User
	err := c.ShouldBindJSON(&u)
	return &u, err
}
