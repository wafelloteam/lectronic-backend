package user

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type user_controller struct {
	service interfaces.UserServiceIF
}

func NewController(service interfaces.UserServiceIF) *user_controller {
	return &user_controller{service}

}

func (c *user_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.GetAll()
	result.Send(w)
}

func (c *user_controller) Add(w http.ResponseWriter, r *http.Request) {

	var data model.User
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

	result := c.service.Add(&data)
	result.Send(w)

}

func (c *user_controller) GetById(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user")
	c.service.GetById(user_id.(string)).Send(w)
}

func (c *user_controller) Update(w http.ResponseWriter, r *http.Request) {

	var data model.User

	UserID := r.Context().Value("user")
	data.UserID = (UserID.(string))

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

	result := c.service.Update(&data)
	result.Send(w)

}

func (c *user_controller) Delete(w http.ResponseWriter, r *http.Request) {

	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.Delete(data.UserID)
	result.Send(w)

}
