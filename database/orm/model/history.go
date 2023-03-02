package model

import "time"

type History struct {
	ID          string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	UserID      string    `gorm:"foreignKey:UserID;references:UserID" json:"user_id" valid:"-"`
	ProductID   string    `json:"product_id" valid:"-"`
	Name        string    `json:"name" valid:"type(string)"`
	Description string    `json:"description" valid:"type(string)"`
	Price       int64     `json:"price" valid:"type(int)"`
	Category    string    `json:"category" valid:"type(string)"`
	ImageURL    string    `json:"image_url" valid:"type(string)"`
	UidImage    string    `json:"uid_image" valid:"type(string)"`
	Qty         int64     `json:"qty" valid:"type(int)"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" valid:"-" `

	User User `json:"user,omitempty"`
}

type Histories []History
