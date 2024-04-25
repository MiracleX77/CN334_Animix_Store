package entities

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		Type      string `json:"type"`
		Status    string `json:"status"`
	}
)
