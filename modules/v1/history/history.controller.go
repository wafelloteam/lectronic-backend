package history

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type history_controller struct {
	service interfaces.HistoryServiceIF
}

func NewController(service interfaces.HistoryServiceIF) *history_controller {
	return &history_controller{service}
}

func (c *history_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user")
	c.service.GetAll(userID.(string)).Send(w)
}

func (c *history_controller) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	c.service.GetById(params).Send(w)
}

func (c *history_controller) AddReview(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var data model.History 
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true)
		return
	}
	// _, err = govalidator.ValidateStruct(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	lib.NewRes(err.Error(), 500, true).Send(w)
	// 	return
	// }

	c.service.AddReview(&data, id).Send(w)
}

func (c *history_controller) GetByProductID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	c.service.GetByProductID(params).Send(w)
}