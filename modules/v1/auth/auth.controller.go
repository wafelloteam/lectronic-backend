package auth

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type auth_controller struct {
	service interfaces.AuthServiceIF
}

func NewController(service interfaces.AuthServiceIF) *auth_controller {
	return &auth_controller{service}

}

func (c *auth_controller) Login(w http.ResponseWriter, r *http.Request) {
	var data model.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true)
		return
	}

	c.service.Login(&data).Send(w)
}

func (c *auth_controller) Register(w http.ResponseWriter, r *http.Request) {
	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true)
		return
	}
	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	c.service.Register(&data).Send(w)

}

func (c *auth_controller) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	var data model.UserPassword
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true)
		return
	}
	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	c.service.ForgetPassword(&data).Send(w)

}

func (c *auth_controller) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)["id"]

	var data model.UserUpdatePassword

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.UpdatePassword(params, &data)
	result.Send(w)

}
