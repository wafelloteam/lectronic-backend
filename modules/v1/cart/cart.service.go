package cart

import (
	"fmt"

	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type cart_service struct {
	repo interfaces.CartRepoIF
}

func NewService(repo interfaces.CartRepoIF) *cart_service {
	return &cart_service{repo}
}

func (s *cart_service) Add(data *model.Cart, productID, userID string) *lib.Response {
	data.UserID = userID
	data.ProductID = productID
	data, err := s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	fmt.Println(data)
	return lib.NewRes(data, 200, false)
}

func (s *cart_service) GetAll(userID string) *lib.Response {
	data, err := s.repo.GetAll(userID)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}


func (s *cart_service) Delete(id string) *lib.Response {
	data, err := s.repo.Delete(id)

	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}


func (s *cart_service) Checkout(data *model.Cart) *lib.Response {
	data, err := s.repo.Checkout(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *cart_service) GetCheckout(userID string) *lib.Response {
	data, err := s.repo.GetCheckout(userID)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *cart_service) Payment() *lib.Response {
	data, err := s.repo.Payment()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
