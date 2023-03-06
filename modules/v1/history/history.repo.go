package history

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type history_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (r *history_repo) GetAll(userID string) (*model.Histories, error) {
	var data model.Histories
	err := r.database.Where("user_id = ?", userID).Preload(clause.Associations).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}


func (r *history_repo) GetById(id string) (*model.History, error) {
	var data model.History

	err := r.database.First(&data, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *history_repo) AddReview(data *model.History, id string) (*model.History, error) {
	var history model.History
	var product model.Product

	tx := r.database.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&history).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}

	err = tx.Model(&history).Where("id = ?", id).Preload(clause.Associations).Find(&history).Error
	if err != nil {
		return nil, err
	}
	
	err = tx.Model(&product).Where("id = ?", history.ProductID).Find(&product).Error
	if err != nil {
		return nil, err
	}
	
	// product.Rating = product.Rating + data.Rating
	// err = tx.Model(&product).Where("id = ?", history.ProductID).Updates(&product).Error
	// if err != nil {
	// 	return nil, err
	// }

	tx.Commit()

	return &history, nil
}

func (r *history_repo) GetByProductID(id string) (*model.ReviewData, error) {
	var history model.History
	var arr model.ReviewData
	
	err := r.database.Model(&history).Where("review != ?", "").Select("review").Find(&arr).Error
	if err != nil {
		return nil, err
	}

	return &arr, nil
}



