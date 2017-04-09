package admin_module

import (
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"net/http"
)

var router = mux.NewRouter()

func withContext(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetContextStore().SetContext(appengine.NewContext(r))
		f(w, r)
	}
}

func init() {

	router.HandleFunc("/countries", withContext(HandleGetAllCountries)).Methods("GET")
	router.HandleFunc("/countries", withContext(HandleSaveCountry)).Methods("POST", "PUT")
	router.HandleFunc("/country/{code}", withContext(HandleGetCountry)).Methods("GET")

	router.HandleFunc("/countries/{countryCode}/universities", withContext(HandleGetAllUniversities)).Methods("GET")
	router.HandleFunc("/countries/{countryCode}/universities", withContext(HandleCreateUniversity)).Methods("POST", "PUT")
	router.HandleFunc("/countries/{countryCode}/universities/{id:[0-9]+}", withContext(HandleGetUniversity)).Methods("GET")

	http.Handle("/", router)
}
