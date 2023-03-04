package auth

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type auth_service struct {
	repo interfaces.UserRepoIF
}

type tokenResponse struct {
	Token string `json:"token"`
}

func NewService(repo interfaces.UserRepoIF) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *model.User) *lib.Response {

	user, err := s.repo.FindEmail(body.Email)
	if err != nil {
		return lib.NewRes("Email not registered", 401, true)
	}

	if lib.CheckPassword(user.Password, body.Password) {
		return lib.NewRes("Wrong password", 401, true)

	}

	jwt := lib.NewToken(user.ID, user.Role)
	token, err := jwt.CreateToken()
	if err != nil {
		return lib.NewRes(err, 501, true)
	}
	return lib.NewRes(tokenResponse{Token: token}, 200, false)

}

func (s *auth_service) Register(body *model.User) *lib.Response {
	_, err := s.repo.FindEmail(body.Email)
	if err == nil {
		return lib.NewRes("Email has been registered", 401, true)
	}
	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	body.Password = hashPassword
	data, err := s.repo.Add(body)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
