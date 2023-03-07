package product

import (
	"errors"
	"fmt"

	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/lib"
	"gorm.io/gorm"
)

type product_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *product_repo {
	return &product_repo{db}

}

func (r *product_repo) Add(data *model.Product) (*model.Product, error) {
	data.Slug = lib.Slug(data.Name)
	err := r.database.Create(data).Error

	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

	return data, nil

}

func (r *product_repo) GetAll() (*model.Products, error) {
	var data model.Products
	err := r.database.Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}

	return &data, nil

}

func (r *product_repo) GetById(id string) (*model.Product, error) {
	var data model.Product

	err := r.database.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

	return &data, nil
}

func (r *product_repo) Update(data *model.Product) (*model.Product, error) {
	err := r.database.Model(&model.Product{}).Where("id = ?", data.ID).Updates(data).Error

	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

	return data, nil
}

func (r *product_repo) Delete(id string) (*model.Product, error) {

	var data model.Product
	result := r.database.Where("id = ?", id).Delete(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil

}

func (r *product_repo) GetByCategory(category string) (*model.Products, error) {
	var data model.Products

	err := r.database.Where("category = ?", category).Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}

	return &data, nil
}

func (r *product_repo) Sort(by string, order string) (*model.Products, error) {
	var data model.Products

	orderQuery := fmt.Sprintf("%s %s", by, order)
	err := r.database.Order(orderQuery).Limit(9).Find(&data).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}

	return &data, nil
}

func (r *product_repo) GetBySlug(slug string) (*model.Product, error) {
	var data model.Product

	err := r.database.First(&data, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	data.Image = lib.ImageReturn(data.Image)

	return &data, nil
}

func (r *product_repo) Search(query string) (*model.Products, error) {

	var data model.Products

	err := r.database.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ? OR LOWER(category) LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&data).Error

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("product not found")
	}

	for i := 0; i < len(data); i++ {
		data[i].Image = lib.ImageReturn(data[i].Image)
	}

	return &data, nil

}
