package admin_module

import (
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
	"net/http"
	"sab.com/domain/country"
	"sab.com/domain/helpers"
	"sab.com/domain/university"
	"strconv"
)

func HandleGetAllUniversities(writer http.ResponseWriter, request *http.Request) {
	universityService := GetUniversityService()

	countryCode := getCountryCodeFromRequest(request)

	if universities, err := universityService.GetAllUniversitiesForCountryCode(countryCode); err != nil {
		log.Errorf(GetContextStore().GetContext(), err.Error())
		writeApiError(writer, []string{"An unknown error occured while getting all universities"}, http.StatusInternalServerError)
	} else {
		writeApiSuccess(writer, universities, http.StatusOK)
	}
}

func HandleCreateUniversity(writer http.ResponseWriter, request *http.Request) {
	universityService := GetUniversityService()
	countryCode := getCountryCodeFromRequest(request)

	var newUniversity university.University

	if err := helpers.JsonToObject(request.Body, &newUniversity); err != nil {
		log.Errorf(GetContextStore().GetContext(), err.Error())
		writeApiError(writer, []string{err.Error()}, http.StatusBadRequest)
		return
	}

	if err := universityService.SaveUniversity(&newUniversity, countryCode); err != nil {
		var errorString string
		if err == country.CountryNotFoundError || err == university.UniversityNotFoundError {
			errorString = err.Error()
		} else {
			errorString = "An unknown error occured while saving the University"
		}
		log.Errorf(GetContextStore().GetContext(), err.Error())
		writeApiError(writer, []string{errorString}, http.StatusInternalServerError)
		return
	}

	writeApiSuccess(writer, newUniversity, http.StatusCreated)
}

func HandleGetUniversity(writer http.ResponseWriter, request *http.Request) {
	universityService := GetUniversityService()
	countryCode, universityId := getCountryCodeAndUniversityIdFromRequest(request)

	if theUniversity, err := universityService.GetUniversityByIdAndCountryCode(universityId, countryCode); err != nil {
		if err == university.UniversityNotFoundError {
			writeApiError(writer, []string{err.Error()}, http.StatusNotFound)
		} else {
			log.Errorf(GetContextStore().GetContext(), err.Error())
			writeApiError(writer, []string{"An unknown error occured while getting country by code"}, http.StatusInternalServerError)
		}
	} else {
		writeApiSuccess(writer, theUniversity, http.StatusOK)
	}
}

func getCountryCodeFromRequest(request *http.Request) string {
	vars := mux.Vars(request)
	return vars["countryCode"]
}

func getCountryCodeAndUniversityIdFromRequest(request *http.Request) (countryCode string, universityId int64) {
	vars := mux.Vars(request)
	countryCode = vars["countryCode"]
	universityId, _ = strconv.ParseInt(vars["id"], 10, 64)
	return
}
