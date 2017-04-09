package admin_module

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
	"net/http"
	"sab.com/domain/helpers"
	"sab.com/domain/university"
	"strconv"
)

func UniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func UniversityHandler(w http.ResponseWriter, r *http.Request) {

}

func HandleGetAllUniversities(writer http.ResponseWriter, request *http.Request) {
	universityService := GetUniversityService()

	countryCode := getCountryCodeFromRequest(request)

	if universities, err := universityService.GetAllUniversitiesForCountryCode(countryCode); err != nil {
		log.Errorf(GetContextStore().GetContext(), "An error occured while getting all universities: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		responseByte, _ := json.Marshal(universities)
		writer.Write(responseByte)
	}
}

func HandleCreateUniversity(writer http.ResponseWriter, request *http.Request) {
	universityService := GetUniversityService()
	countryCode := getCountryCodeFromRequest(request)

	var newUniversity university.University

	if err := helpers.JsonToObject(request.Body, &newUniversity); err != nil {
		log.Errorf(GetContextStore().GetContext(), "Could not convert request body to University: %s", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := universityService.SaveUniversity(&newUniversity, countryCode); err != nil {
		log.Errorf(GetContextStore().GetContext(), "An error occured while saving the University: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseByte, _ := json.Marshal(newUniversity)
	writer.Write(responseByte)
	writer.WriteHeader(http.StatusCreated)
}

func HandleGetUniversity(writer http.ResponseWriter, request *http.Request) {
	universityService := GetUniversityService()
	countryCode, universityId := getCountryCodeAndUniversityIdFromRequest(request)

	if theUniversity, err := universityService.GetUniversityByIdAndCountryCode(universityId, countryCode); err != nil {
		if err == university.UniversityNotFoundError {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprint(writer, err)
		} else {
			log.Errorf(GetContextStore().GetContext(), "An error occured while getting country by code: %s", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		responseByte, _ := json.Marshal(theUniversity)
		writer.Write(responseByte)
		writer.WriteHeader(http.StatusOK)
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
