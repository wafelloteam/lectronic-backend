package interfaces

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type HistoryRepoIF interface {
	GetAll(userID string) (*model.Histories, error)
	GetById(id string) (*model.History, error)
	AddReview(data *model.History, id string) (*model.History, error)
	GetByProductID(id string) (*model.ReviewData, error)
}

type HistoryServiceIF interface {
	GetAll(userID string) *lib.Response
	GetById(id string) *lib.Response 
	AddReview(data *model.History, id string) *lib.Response
	GetByProductID(id string) *lib.Response 
}
