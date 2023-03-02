package user

import (
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/id", middleware.Handle(controller.GetById, middleware.AuthMiddleware("user"))).Methods("GET")
	route.HandleFunc("/all", middleware.Handle(controller.GetAll, middleware.AuthMiddleware("admin"))).Methods("GET")
	route.HandleFunc("/add", middleware.Handle(controller.Add, middleware.AuthMiddleware("user"))).Methods("POST")
	route.HandleFunc("/update", middleware.Handle(controller.Update, middleware.AuthMiddleware("user"))).Methods("PUT")
	route.HandleFunc("/delete", middleware.Handle(controller.Delete, middleware.AuthMiddleware("admin"))).Methods("DELETE")

}
