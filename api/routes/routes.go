package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	routesModel "go-lang-web-service/models/routes"
	routesService "go-lang-web-service/services/routes"
	"net/http"
	"net/url"
)

// Returns the list of Route's destances and durations starting from the given source
func GetRoutes(w http.ResponseWriter, r *http.Request) {
	source, dists, err := parseQuery(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := routesModel.AllRoutes{
		Source: source,
		Routes: make([]routesModel.RouteInfo, 0, len(dists)),
	}
	for _, dist := range dists {
		routes, err := routesService.GetAllRoutes(source, dist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result.Routes = append(result.Routes, routes.Routes...)
	}
	fmt.Fprint(w, result)
	json.NewEncoder(w).Encode(result)
}

// Parse the Query Params map and returns the source and distinations
func parseQuery(values url.Values) (source string, dists []string, err error) {
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
