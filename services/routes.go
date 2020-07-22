package services

import (
	"github.com/amine-bambrik-p8/go-lang-web-service/common/async"
	"github.com/amine-bambrik-p8/go-lang-web-service/common/http"

	"log"
	"net/url"
	"path"

	"github.com/amine-bambrik-p8/go-lang-web-service/models"
)

const BaseURL = "http://router.project-osrm.org/route/v1/driving/"

type IRoutesService interface {
	GetAllRoutes(source string, dists []string) (allRoutes *models.AllRoutes, err error)
	GetRoutes(source string, dist string) (routes *models.AllRoutes, err error)
}
type RoutesService struct {
}

var (
	Routes IRoutesService
)

func init() {
	Routes = &RoutesService{}
}

// Returns a list of all routes from a source to multiple possible destinations
func (r *RoutesService) GetAllRoutes(source string, dists []string) (allRoutes *models.AllRoutes, err error) {
	requests := make([]async.Promise, 0, len(dists))
	for _, dist := range dists {
		promise := async.NewPromise(func() (interface{}, error) { return r.GetRoutes(source, dist) })
		requests = append(requests, promise)
	}
	promise := async.WaitAll(requests...)

	routes, err := promise.Await()
	if err != nil {
		return
	}

	allRoutes = &models.AllRoutes{
		Source: source,
		Routes: make([]models.RouteInfo, 0, len(dists)),
	}
	for _, route := range routes.([]interface{}) {
		obj := route.(*models.AllRoutes)
		allRoutes.Routes = append(allRoutes.Routes, obj.Routes...)
	}
	return
}

// Return a list of all possible routes from a source to destination
func (r *RoutesService) GetRoutes(source string, dist string) (routes *models.AllRoutes, err error) {
	endpoint := getURL(source, dist)
	err = http.GetRequestJSON(endpoint, &routes)
	if err != nil {
		return
	}
	routes.Source = source
	for idx, _ := range routes.Routes {
		routes.Routes[idx].Destination = dist
	}
	return
}

func getURL(source string, dist string) (endpoint string) {
	endpointURL, err := url.Parse(BaseURL)
	if err != nil {
		log.Print("Invalid Base URL !!!")

		panic("Invalid Base URL !!!")
	}
	endpointURL.Scheme = "http"
	endpointURL.Path = path.Join(endpointURL.Path, source+";"+dist)
	values := endpointURL.Query()
	values.Set("overview", "false")
	endpointURL.RawQuery = values.Encode()
	endpoint = endpointURL.String()
	return
}
