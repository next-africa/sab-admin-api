package admin_module

import (
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func init() {

	router.HandleFunc("/countries", CountriesHandler).Methods("GET", "POST")
	router.HandleFunc("/country/{code}", CountryHandler).Methods("GET")

	router.HandleFunc("/universities", UniversitiesHandler)
	router.HandleFunc("/university/{id}", UniversityHandler)

	http.Handle("/", router)
}
