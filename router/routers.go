package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wafellofazztrack/lectronic-backend/database/orm"
	"github.com/wafellofazztrack/lectronic-backend/modules/v1/auth"
	"github.com/wafellofazztrack/lectronic-backend/modules/v1/cart"
	"github.com/wafellofazztrack/lectronic-backend/modules/v1/history"
	"github.com/wafellofazztrack/lectronic-backend/modules/v1/product"
	"github.com/wafellofazztrack/lectronic-backend/modules/v1/user"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.NewDB()
	if err != nil {
		return nil, err
	}

	var imageFolder = http.FileServer(http.Dir("./public/image"))
	mainRoute.PathPrefix("/public/").Handler(http.StripPrefix("/public/image", imageFolder))

	user.NewRoute(mainRoute, db)
	auth.NewRoute(mainRoute, db)
	product.NewRoute(mainRoute, db)
	cart.NewRoute(mainRoute, db)
	history.NewRoute(mainRoute, db)

	mainRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("Hello World! This is lectronic-api. You can check Postman Documentation <a href=\"https://documenter.getpostman.com/view/25042327/2s93JtQPYk\">here</a>"))
	})

	return mainRoute, nil

}
