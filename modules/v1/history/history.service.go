package history

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type history_service struct {
	repo interfaces.HistoryRepoIF
}

func NewService(repo interfaces.HistoryRepoIF) *history_service {
	return &history_service{repo}
}

func (s *history_service) GetAll(userID string) *lib.Response {
	data, err := s.repo.GetAll(userID)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *history_service) GetById(id string) *lib.Response {
	data, err := s.repo.GetById(id)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *history_service) AddReview(data *model.History, id string) *lib.Response {
	data, err := s.repo.AddReview(data, id)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *history_service) GetByProductID(id string) *lib.Response {
	data, err := s.repo.GetByProductID(id)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}
