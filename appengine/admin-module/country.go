package admin_module

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
	"net/http"
	"sab.com/domain/country"
	"sab.com/domain/helpers"
)

func HandleGetAllCountries(writer http.ResponseWriter, request *http.Request) {

	countryService := GetCountryService()

	if countries, err := countryService.GetAllCountries(); err != nil {
		log.Errorf(GetContextStore().GetContext(), "An error occured while getting all countries: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		responseByte, _ := json.Marshal(countries)
		writer.Write(responseByte)
	}
}

func HandleSaveCountry(writer http.ResponseWriter, request *http.Request) {
	countryService := GetCountryService()

	var newCountry country.Country

	if err := helpers.JsonToObject(request.Body, &newCountry); err != nil {
		log.Errorf(GetContextStore().GetContext(), "Could not convert request body to Country: %s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := countryService.SaveCountry(&newCountry); err != nil {
		log.Errorf(GetContextStore().GetContext(), "An error occured while saving the Country: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseByte, _ := json.Marshal(newCountry)
	writer.Write(responseByte)
	writer.WriteHeader(http.StatusCreated)
}

func HandleGetCountry(w http.ResponseWriter, r *http.Request) {
	countryService := GetCountryService()

	vars := mux.Vars(r)
	code := vars["code"]

	if theCountry, err := countryService.GetCountryByCode(code); err != nil {
		if err == country.CountryNotFoundError {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, err)
		} else {
			log.Errorf(GetContextStore().GetContext(), "An error occured while getting country by code: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		responseByte, _ := json.Marshal(theCountry)
		w.Write(responseByte)
		w.WriteHeader(http.StatusOK)
	}
}
