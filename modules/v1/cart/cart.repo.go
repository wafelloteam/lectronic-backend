package cart

import (
	"github.com/google/uuid"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cart_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *cart_repo {
	return &cart_repo{db}

}

func (r *cart_repo) Add(data *model.Cart) (*model.Cart, error) {
	data.IsChecked = false
	data.Qty = 1
	err := r.database.Create(data).Error
	if err != nil {
		return nil, err
	}

	r.database.Where("id = ?", data.ID).Preload(clause.Associations).First(&data)
	data.User.Password = ""

	return data, nil

}

func (r *cart_repo) GetAll(userID string) (*model.Carts, error) {
	var data model.Carts
	err := r.database.Where("user_id = ?", userID).Preload(clause.Associations).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *cart_repo) Delete(id string) (*model.Cart, error) {
	var data model.Cart
	result := r.database.Where("id = ?", id).Delete(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil

}

func (r *cart_repo) Checkout(data *model.Cart) (*model.Cart, error) {
	data.IsChecked = true
	err := r.database.Model(&model.Cart{}).Where("id = ?", data.ID).Preload(clause.Associations).Updates(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *cart_repo) GetCheckout(userID string) (*model.Carts, error) {
	var data model.Carts
	err := r.database.Where("user_id = ?", userID).Where("is_checked = true").Preload(clause.Associations).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *cart_repo) Payment() (*model.Histories, error) {

	var carts model.Carts
	var product model.Product
	var history model.History
	var histories model.Histories

	tx := r.database.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&carts).Where("is_checked = true").Preload(clause.Associations).Find(&carts).Error
	if err != nil {
		return nil, err
	}

	for _, cart := range carts {
		history.ID = uuid.NewString()
		history.UserID = cart.UserID
		history.ProductID = cart.ProductID
		history.Name = cart.Product.Name
		history.Description = cart.Product.Description
		history.Price = cart.Product.Price
		history.Category = cart.Product.Category
		history.ImageURL = cart.Product.ImageURL
		history.UidImage = cart.Product.UidImage
		history.Qty = cart.Qty

		err = tx.Model(&history).Create(&history).Error
		if err != nil {
			return nil, err
		}

		product.Sold = cart.Product.Sold + 1
		product.Stock = cart.Product.Stock - cart.Qty
		err = tx.Model(&product).Where("id = ?", cart.ProductID).Updates(&product).Error
		if err != nil {
			return nil, err
		}

		err = tx.Model(&cart).Where("is_checked = true").Delete(&cart).Error
		if err != nil {
			return nil, err
		}
	}

	err = tx.Model(&histories).Preload(clause.Associations).Find(&histories).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &histories, nil
}
