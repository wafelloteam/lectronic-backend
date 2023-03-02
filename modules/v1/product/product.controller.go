package product

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/interfaces"
	"github.com/wafellofazztrack/lectronic-backend/lib"
)

type product_controller struct {
	service interfaces.ProductServiceIF
}

func NewController(service interfaces.ProductServiceIF) *product_controller {
	return &product_controller{service}

}

func (c *product_controller) Add(w http.ResponseWriter, r *http.Request) {

	var data model.Product

	image := r.Context().Value("image").(string)
	data.Image = image

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(image)
		lib.NewRes(err.Error(), 400, true).Send(w)
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

func (c *product_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.GetAll()
	result.Send(w)
}

func (c *product_controller) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	c.service.GetById(params).Send(w)
}

func (c *product_controller) Update(w http.ResponseWriter, r *http.Request) {

	var data model.Product

	image := r.Context().Value("image").(string)
	data.Image = image

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(image)
		lib.NewRes(err.Error(), 400, true).Send(w)
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

func (c *product_controller) Delete(w http.ResponseWriter, r *http.Request) {

	var data model.Product
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.Delete(data.ID)
	result.Send(w)

}

func (c *product_controller) GetByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["category"]
	c.service.GetByCategory(params).Send(w)
}

func (c *product_controller) Sort(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	by := query["by"][0]
	order := query["order"][0]
	c.service.Sort(by, order).Send(w)
}

func (c *product_controller) GetBySlug(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["slug"]
	c.service.GetBySlug(params).Send(w)
}

func (c *product_controller) Search(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	s := query["s"][0]

	c.service.Search(s).Send(w)

}
