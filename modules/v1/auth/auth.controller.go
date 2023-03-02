package auth

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
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
