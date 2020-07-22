package routes

import (
	"errors"
	"go-lang-web-service/common"
	routesModel "go-lang-web-service/models/routes"
	routesService "go-lang-web-service/services/routes"

	"net/http"
	"net/url"
)

// Controller Object for Routes model routes(endpoints)
var Controller = &RoutesController{}

type RoutesController struct {
	common.Controller
}

// Returns the list of Route's destances and durations starting from the given source
func (c *RoutesController) GetRoutes(w http.ResponseWriter, r *http.Request) {
	source, dists, err := c.parseQuery(r.URL.Query())
	if err != nil {
		c.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}

	result := routesModel.AllRoutes{
		Source: source,
		Routes: make([]routesModel.RouteInfo, 0, len(dists)),
	}
	for _, dist := range dists {
		routes, err := routesService.GetAllRoutes(source, dist)
		if err != nil {
			c.SendJSON(w, r, err, http.StatusBadRequest)
			return
		}
		result.Routes = append(result.Routes, routes.Routes...)
	}
	result.SortByDuration()
	c.SendJSON(w, r, result.GetViewModel(), http.StatusOK)
}

// Parse the Query Params map and returns the source and distinations
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
