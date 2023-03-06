package history

import (
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/all", middleware.Handle(controller.GetAll, middleware.AuthMiddleware("user"))).Methods("GET")
	route.HandleFunc("/id/{id}", middleware.Handle(controller.GetById, middleware.AuthMiddleware("user"))).Methods("GET")
	route.HandleFunc("/addReview/{id}", middleware.Handle(controller.AddReview, middleware.AuthMiddleware("user"))).Methods("PUT")
	route.HandleFunc("/review/{id}",controller.GetByProductID).Methods("GET")

}
