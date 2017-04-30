package admin_module

import (
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
	"net/http"
	"sab.com/domain/country"
	"sab.com/domain/helpers"
)

func HandleGetAllCountries(writer http.ResponseWriter, request *http.Request) {
	countryService := GetCountryService()

	if countries, err := countryService.GetAllCountries(); err != nil {
		log.Errorf(GetContextStore().GetContext(), err.Error())
		writeApiError(writer, []string{"An unknow error occured while getting all countries"}, http.StatusInternalServerError)
	} else {
		writeApiSuccess(writer, countries, http.StatusOK)
	}
}

func HandleSaveCountry(writer http.ResponseWriter, request *http.Request) {
	countryService := GetCountryService()

	var newCountry country.Country

	if err := helpers.JsonToObject(request.Body, &newCountry); err != nil {
		log.Errorf(GetContextStore().GetContext(), err.Error())
		writeApiError(writer, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if err := countryService.SaveCountry(&newCountry); err != nil {
		log.Errorf(GetContextStore().GetContext(), err.Error())
		writeApiError(writer, []string{"An error occured while saving the Country"}, http.StatusInternalServerError)
		return
	}

	writeApiSuccess(writer, newCountry, http.StatusCreated)
}

func HandleGetCountry(w http.ResponseWriter, r *http.Request) {
	countryService := GetCountryService()

	vars := mux.Vars(r)
	code := vars["code"]

	if theCountry, err := countryService.GetCountryByCode(code); err != nil {
		if err == country.CountryNotFoundError {
			writeApiError(w, []string{err.Error()}, http.StatusNotFound)
		} else {
			log.Errorf(GetContextStore().GetContext(), err.Error())
			writeApiError(w, []string{"An unknow error occured while getting country by code"}, http.StatusInternalServerError)
		}
	} else {
		writeApiSuccess(w, theCountry, http.StatusOK)
	}
}
