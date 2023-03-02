package user

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type user_service struct {
	repo interfaces.UserRepoIF
}

func NewService(repo interfaces.UserRepoIF) *user_service {
	return &user_service{repo}

}

func (s *user_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *user_service) Add(data *model.User) *lib.Response {
	hashPassword, err := lib.HashPassword(data.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	data.Password = hashPassword
	data, err = s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *user_service) GetById(uuid string) *lib.Response {
	data, err := s.repo.GetById(uuid)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *user_service) Update(body *model.User) *lib.Response {

	_, err := s.repo.FindEmail(body.Email)
	if err == nil {
		return lib.NewRes("Email has been registered", 401, true)
	}
	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	body.Password = hashPassword

	data, err := s.repo.Update(body)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *user_service) Delete(uuid string) *lib.Response {
	data, err := s.repo.Delete(uuid)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
