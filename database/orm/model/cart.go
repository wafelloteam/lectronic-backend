package model

import "time"

type Cart struct {
	ID        string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	UserID    string    `gorm:"foreignKey:ID;references:ID" json:"user_id" valid:"-"`
	ProductID string    `gorm:"foreignKey:ID;references:ID" json:"product_id" valid:"-"`
	Qty       int64     `gorm:"default:1" json:"qty" valid:"type(int)"`
	IsChecked bool      `gorm:"default:true" json:"is_checked"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty" valid:"-" `

	User    User    `json:"user,omitempty"`
	Product Product `json:"product,omitempty"`
}

type Carts []Cart
