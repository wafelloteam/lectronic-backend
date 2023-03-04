package cart

import (
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/cart").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/add/{id}", middleware.Handle(controller.Add, middleware.AuthMiddleware("user"))).Methods("POST")
	route.HandleFunc("/all", middleware.Handle(controller.GetAll, middleware.AuthMiddleware("user"))).Methods("GET")
	route.HandleFunc("/delete", middleware.Handle(controller.Delete, middleware.AuthMiddleware("user"))).Methods("DELETE")

	route.HandleFunc("/checkout", middleware.Handle(controller.GetCheckout, middleware.AuthMiddleware("user"))).Methods("GET")
	route.HandleFunc("/checkout", middleware.Handle(controller.Checkout, middleware.AuthMiddleware("user"))).Methods("PUT")
	
	route.HandleFunc("/payment", middleware.Handle(controller.Payment, middleware.AuthMiddleware("user"))).Methods("POST")
}
