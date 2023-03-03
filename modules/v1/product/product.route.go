package product

import (
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/add", middleware.Handle(controller.Add, middleware.AuthMiddleware("admin"), middleware.AuthUploadImage())).Methods("POST")
	route.HandleFunc("/update", middleware.Handle(controller.Update, middleware.AuthMiddleware("admin"), middleware.AuthUploadImage())).Methods("PUT")
	route.HandleFunc("/delete", middleware.Handle(controller.Delete, middleware.AuthMiddleware("admin"))).Methods("DELETE")
	route.HandleFunc("/all", controller.GetAll).Methods("GET")
	route.HandleFunc("/id/{id}", controller.GetById).Methods("GET")
	route.HandleFunc("/category/{category}", controller.GetByCategory).Methods("GET")
	route.HandleFunc("/sort", controller.Sort).Methods("GET")
	route.HandleFunc("/slug/{slug}", controller.GetBySlug).Methods("GET")
	route.HandleFunc("/search", controller.Search).Methods("GET")

}
