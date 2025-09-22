package models

// User represents a user in the system
// @Description User entity with authentication and role information
type User struct {
	Base
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" gorm:"unique" example:"john.doe@example.com"`
	Password string `json:"-"`
	Role     string `json:"role" example:"admin"`
}
