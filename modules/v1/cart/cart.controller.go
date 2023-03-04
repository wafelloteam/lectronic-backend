package cart

import (
	"encoding/json"
	// "fmt"
	"net/http"

	// "github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type cart_controller struct {
	service interfaces.CartServiceIF
}

func NewController(service interfaces.CartServiceIF) *cart_controller {
	return &cart_controller{service}
}

func (c *cart_controller) Add(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user")
	productID := mux.Vars(r)["id"]
	var data model.Cart 
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

	c.service.Add(&data, productID, userID.(string)).Send(w)
}

func (c *cart_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user")
	c.service.GetAll(userID.(string)).Send(w)
}

func (c *cart_controller) Delete(w http.ResponseWriter, r *http.Request) {

	var data model.Cart
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.Delete(data.ID)
	result.Send(w)

}


func (c *cart_controller) GetCheckout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user")
	c.service.GetCheckout(userID.(string)).Send(w)
}


func (c *cart_controller) Checkout(w http.ResponseWriter, r *http.Request) {
	var data model.Cart 
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

	c.service.Checkout(&data).Send(w)
}

func (c *cart_controller) Payment(w http.ResponseWriter, r *http.Request) {
	c.service.Payment().Send(w)
}
