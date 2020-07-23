package handlers

import (
	"errors"
	"net/http"
	"net/url"

	commonHTTP "github.com/amine-bambrik-p8/go-lang-web-service/common/http"
	"github.com/amine-bambrik-p8/go-lang-web-service/services"
	"github.com/gorilla/mux"
)

// Hooks up all the api's routes for the RoutesController
func (c *RoutesController) HookHandlers() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/routes", c.GetRoutes).Methods("GET")
	return router
}

// IRoutesController interface of OSRM Routes Handler
type IRoutesController interface {
	GetRoutes(w http.ResponseWriter, r *http.Request)
	HookHandlers() http.Handler
}

// RoutesController struct that represents the OSRM Routes handler
type RoutesController struct {
}

// the OSRM Routes handler
// NOTE: the RoutesHandler can be easily mocked since it's made public as var
var (
	RoutesHandler IRoutesController
)

func init() {
	RoutesHandler = &RoutesController{}
}

// Returns the list of Route's destances and durations starting from the given source
// TODO should implement validation for src and dst QueryParms
func (c *RoutesController) GetRoutes(w http.ResponseWriter, r *http.Request) {
	source, dists, err := c.parseQuery(r.URL.Query())
	if err != nil {
		commonHTTP.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}

	allRoutes, err := services.Routes.GetAllRoutes(source, dists)
	if err != nil {
		commonHTTP.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}
	allRoutes.SortByDuration()
	commonHTTP.SendJSON(w, r, allRoutes, http.StatusOK)
}

// Parse the Query Params map and returns the source and destinations
// TODO should use mux in the future for QueryParams
func (c RoutesController) parseQuery(values url.Values) (source string, dists []string, err error) {
	src, ok := values["src"]
	if !ok || len(src) != 1 {
		err = errors.New("Missing 'src' URL param")
		return
	}
	source = src[0]
	dists, ok = values["dst"]
	if !ok || len(dists) < 1 {
		err = errors.New("Missing 'dst' URL param")
		return
	}
	return
}
