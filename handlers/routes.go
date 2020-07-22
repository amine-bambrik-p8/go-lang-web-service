package handlers

import (
	"errors"
	"go-lang-web-service/common/handlers"
	"go-lang-web-service/services"
	"net/http"
	"net/url"
)

// Controller Object for Routes model routes(endpoints)
var RoutesHandler = &RoutesController{}

type RoutesController struct {
	handlers.Controller
}

// Returns the list of Route's destances and durations starting from the given source
func (c *RoutesController) GetRoutes(w http.ResponseWriter, r *http.Request) {
	source, dists, err := c.parseQuery(r.URL.Query())
	if err != nil {
		c.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}

	allRoutes, err := services.Routes.GetAllRoutes(source, dists)
	if err != nil {
		c.SendJSON(w, r, err, http.StatusBadRequest)
		return
	}
	allRoutes.SortByDuration()
	c.SendJSON(w, r, allRoutes, http.StatusOK)
}

// Parse the Query Params map and returns the source and destinations
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
