package product

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type product_service struct {
	repo interfaces.ProductRepoIF
}

func NewService(repo interfaces.ProductRepoIF) *product_service {
	return &product_service{repo}

}

func (s *product_service) Add(data *model.Product) *lib.Response {
	data, err := s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *product_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *product_service) GetById(id string) *lib.Response {
	data, err := s.repo.GetById(id)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *product_service) Update(data *model.Product) *lib.Response {
	data, err := s.repo.Update(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *product_service) Delete(id string) *lib.Response {
	data, err := s.repo.Delete(id)

	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *product_service) GetByCategory(category string) *lib.Response {
	data, err := s.repo.GetByCategory(category)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *product_service) Sort(by string, order string) *lib.Response {
	data, err := s.repo.Sort(by, order)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *product_service) GetBySlug(slug string) *lib.Response {
	data, err := s.repo.GetBySlug(slug)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *product_service) Search(query string) *lib.Response {

	data, err := s.repo.Search(query)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)

}
