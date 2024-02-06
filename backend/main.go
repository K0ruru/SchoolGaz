// main.go

package main

import (
	"log"
	"net/http"
	"server/db"
	"server/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.InitializeDB()

	router := mux.NewRouter()

	corsHandler := cors.Default().Handler(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	routes.RouteUser(router)

	port := ":8080"

	log.Printf("Server started on http://localhost%s", port)

	err := http.ListenAndServe(port, corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}
