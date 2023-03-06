package interfaces

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type UserRepoIF interface {
	GetAll() (*model.Users, error)
	Add(data *model.User) (*model.User, error)
	FindEmail(email string) (*model.User, error)
	GetById(uuid string) (*model.User, error)
	Update(data *model.User) (*model.User, error)
	Delete(uuid string) (*model.User, error)
	UpdatePassword(id string, data *model.UserUpdatePassword) (*model.UserUpdatePassword, error)
}

type UserServiceIF interface {
	GetAll() *lib.Response
	Add(data *model.User) *lib.Response
	GetById(uuid string) *lib.Response
	Update(body *model.User) *lib.Response
	Delete(uuid string) *lib.Response
}
