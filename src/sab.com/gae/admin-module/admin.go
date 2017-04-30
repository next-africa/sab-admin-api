package admin_module

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"net/http"
	"sab.com/api"
	"sab.com/graphql"
)

var router = mux.NewRouter()

func withContext(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetContextStore().SetContext(appengine.NewContext(r))
		w.Header().Add("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}

func writeApiError(writer http.ResponseWriter, errors interface{}, status int) {
	apiResponse := api.ApiResponse{Errors: errors}
	writeApiResponse(writer, apiResponse, status)
}

func writeApiSuccess(writer http.ResponseWriter, data interface{}, status int) {
	apiResponse := api.ApiResponse{Data: data}
	writeApiResponse(writer, apiResponse, status)
}

func writeApiResponse(writer http.ResponseWriter, response api.ApiResponse, status int) {
	responseByte, _ := json.Marshal(response)
	writer.WriteHeader(status)
	writer.Write(responseByte)
}

type AppengineGraphqlHandler struct {
	h *graphql.GraphqlHandler
}

func (graphqlHandler AppengineGraphqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	withContext(graphqlHandler.h.ServeHTTP)(w, r)
}

func init() {

	router.HandleFunc("/api/countries", withContext(HandleGetAllCountries)).Methods("GET")
	router.HandleFunc("/api/countries", withContext(HandleSaveCountry)).Methods("POST", "PUT")
	router.HandleFunc("/api/countries/{code}", withContext(HandleGetCountry)).Methods("GET")

	router.HandleFunc("/api/countries/{countryCode}/universities", withContext(HandleGetAllUniversities)).Methods("GET")
	router.HandleFunc("/api/countries/{countryCode}/universities", withContext(HandleCreateUniversity)).Methods("POST", "PUT")
	router.HandleFunc("/api/countries/{countryCode}/universities/{id:[0-9]+}", withContext(HandleGetUniversity)).Methods("GET")

	http.Handle("/api/", router)
	http.Handle("/graphql", AppengineGraphqlHandler{graphql.GetGraphqlHandler(GetCountryService(), GetUniversityService())})
}
