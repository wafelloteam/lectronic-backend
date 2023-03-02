package user

import (
	"errors"

	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"gorm.io/gorm"
)

type user_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}

}

func (r *user_repo) GetById(id string) (*model.User, error) {

	var data model.User

	result := r.database.First(&data, "id = ?", id)
	data.Password = ""

	if result.Error != nil {
		return nil, errors.New("user does not exist")
	}

	return &data, nil
}

func (r *user_repo) GetAll() (*model.Users, error) {
	var data model.Users
	err := r.database.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *user_repo) Add(data *model.User) (*model.User, error) {
	err := r.database.Create(data).Error
	if err != nil {
		return nil, err
	}
	data.Password = ""

	return data, nil

}

func (r *user_repo) Update(data *model.User) (*model.User, error) {
	err := r.database.Model(&model.User{}).Where("id = ?", data.ID).Updates(data).Error

	if err != nil {
		return nil, err
	}
	data.Password = ""

	return data, nil
}

func (r *user_repo) Delete(id string) (*model.User, error) {
	var data model.User
	result := r.database.Where("id = ?", id).Delete(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil

}

func (r *user_repo) FindEmail(email string) (*model.User, error) {
	var data model.User

	result := r.database.First(&data, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("email not found")
	}

	return &data, nil
}
