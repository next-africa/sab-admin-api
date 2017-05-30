package admin_module

import (
	"google.golang.org/appengine"
	"net/http"
	"sab.com/graphql"
)

func withContext(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetContextStore().SetContext(appengine.NewContext(r))
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept")
		f(w, r)
	}
}

type AppengineGraphqlHandler struct {
	h *graphql.GraphqlHandler
}

func (graphqlHandler AppengineGraphqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	withContext(graphqlHandler.h.ServeHTTP)(w, r)
}

func init() {
	http.Handle("/graphql", AppengineGraphqlHandler{graphql.GetGraphqlHandler(GetCountryService(), GetUniversityService())})
}
