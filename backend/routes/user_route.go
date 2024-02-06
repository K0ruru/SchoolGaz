// routes/user_routes.go

package routes

import (
	"server/handlers"

	"github.com/gorilla/mux"
)

func RouteUser(r *mux.Router) {
	r.HandleFunc("/user/{nis}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/add", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")
}
