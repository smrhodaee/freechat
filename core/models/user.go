package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRole uint8

const (
	UserRoleNormal UserRole = iota
	UserRoleOfficial
	UserRoleStaff
	UserRoleSuperAdmin
)

type User struct {
	Username  string    `gorm:"primary_key" json:"username"`
	Password  string    `json:"password"`
	Role      UserRole  `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) IsMatchedPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (User) Table() string {
	return "user"
}
