package admin_module

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"net/http"
	"sab.com/countrystore"
	"sab.com/helpers"
)

func CountriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handleCreateCountry(w, r)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func handleCreateCountry(writer http.ResponseWriter, request *http.Request) {
	var country countrystore.Country
	helpers.JsonToObject(request.Body, &country)

	countrystore.SaveCountry(country, appengine.NewContext(request))
	writer.WriteHeader(http.StatusOK)
}

func CountryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	country, err := countrystore.GetCountryByCode(code, appengine.NewContext(r))

	if err != nil {
		if _, ok := err.(*countrystore.CountryNotFoundError); ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, err)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
	}

	responseByte, _ := json.Marshal(country)

	w.Write(responseByte)
	w.WriteHeader(http.StatusOK)
}
