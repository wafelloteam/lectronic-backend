package model

import (
	"time"
)

type Product struct {
	ID          string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name        string    `json:"name" valid:"type(string)"`
	Description string    `json:"description" valid:"type(string)"`
	Price       int64     `json:"price" valid:"type(int64)"`
	Category    string    `json:"category" valid:"type(string)"`
	Rating      int64     `gorm:"default:0" json:"rating,omitempty" valid:"-"`
	Stock       int64     `gorm:"default:1" json:"stock,omitempty" valid:"-"`
	Sold        int64     `gorm:"default:0" json:"sold,omitempty" valid:"-"`
	Image       string    `json:"image,omitempty" valid:"-"`
	ImageURL    string    `json:"image_url,omitempty" valid:"-"`
	UidImage    string    `json:"uid_image,omitempty" valid:"-"`
	Slug        string    `json:"slug,omitempty" valid:"-"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" valid:"-" `
}

type Products []Product
