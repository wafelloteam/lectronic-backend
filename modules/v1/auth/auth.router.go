package auth

import (
	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/modules/v1/user"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := user.NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/login", controller.Login).Methods("POST")
	route.HandleFunc("/register", controller.Register).Methods("POST")
	route.HandleFunc("/forget-password", controller.ForgetPassword).Methods("POST")
	route.HandleFunc("/update-password/{id}", controller.UpdatePassword).Methods("PUT")

}
