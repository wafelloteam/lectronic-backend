package interfaces

import (
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type AuthServiceIF interface {
	Login(user *model.User) *lib.Response
	Register(user *model.User) *lib.Response
	ForgetPassword(body *model.UserPassword) *lib.Response
	UpdatePassword(id string, body *model.UserUpdatePassword) *lib.Response
}
