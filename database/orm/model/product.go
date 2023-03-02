package model

import (
	"time"
)

type Product struct {
	ID          string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"user_id,omitempty" valid:"-"`
	Name        string    `json:"name" valid:"type(string)"`
	Description string    `json:"description" valid:"type(string)"`
	Price       int64     `json:"price" valid:"type(int)"`
	Category    string    `json:"category" valid:"type(string)"`
	Rating      int64     `json:"rating" valid:"type(int)"`
	Stock       int64     `json:"stock" valid:"type(int)"`
	Sold        int64     `json:"sold" valid:"type(int)"`
	ImageURL    string    `json:"image_url" valid:"type(string)"`
	UidImage    string    `json:"uid_image" valid:"type(string)"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" valid:"-" `
}

type Products []Product
