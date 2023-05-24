package routes

import (
	"crud_mux/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!\n"))
	})

	r.HandleFunc("/nationality", controllers.NationalityIndex).Methods("GET")
	r.HandleFunc("/nationality", controllers.NationalityAdd).Methods("POST")
	r.HandleFunc("/nationality", controllers.NationalityUpdate).Methods("PUT")
	r.HandleFunc("/nationality", controllers.NationalityDelete).Methods("DELETE")

	r.HandleFunc("/customer", controllers.CustomerIndex).Methods("GET")
	r.HandleFunc("/customer", controllers.CustomerAdd).Methods("POST")
	r.HandleFunc("/customer", controllers.CustomerUpdate).Methods("PUT")
	r.HandleFunc("/customer", controllers.CustomerDelete).Methods("DELETE")

	r.HandleFunc("/family-list", controllers.FamilyIndex).Methods("GET")
	r.HandleFunc("/family-list", controllers.FamilyAdd).Methods("POST")
	r.HandleFunc("/family-list", controllers.FamilyUpdate).Methods("PUT")
	r.HandleFunc("/family-list", controllers.FamilyDelete).Methods("Delete")
}
