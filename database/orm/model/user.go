package model

import (
	"time"
)

type User struct {
	ID          string     `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	FullName    string     `json:"full_name" valid:"type(string)"`
	Email       string     `json:"email" valid:"email"`
	Password    string     `json:"password,omitempty" valid:"type(string)"`
	Gender      string     `json:"gender,,omitempty" valid:"-"`
	Address     string     `json:"address,omitempty" valid:"-"`
	Birthday    *time.Time `json:"birthday,omitempty" valid:"-"`
	PhoneNumber string     `json:"phone_number,omitempty" valid:"-"`
	Role        string     `gorm:"default:user" json:"role,omitempty" valid:"-"`
	CreatedAt   time.Time  `json:"created_at" valid:"-"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" valid:"-" `
}

type Users []User
