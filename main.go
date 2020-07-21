package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

type RouteInfo struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

type AllRoutes struct {
	Source string      `json:"source"`
	Routes []RouteInfo `json:"routes"`
}

const BaseURL = "http://router.project-osrm.org/route/v1/driving/"

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

/**
 */

func getAllRoutes(source string, dist string) (routes *AllRoutes, err error) {
	endpoint := getURL(source, dist)
	res, err := http.Get(endpoint)
	if err != nil {
		log.Print("Failed to request routes")
		return
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Print("Failed to parse request body")
		return
	}

	if err = json.Unmarshal(bytes, &routes); err != nil {
		fmt.Println("Error parsing json", err)
		return
	}
	routes.Source = source
	for idx, _ := range routes.Routes {
		routes.Routes[idx].Destination = dist
	}
	return
}
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
func getRoutes(w http.ResponseWriter, r *http.Request) {
	source, dists, err := parseQuery(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result := AllRoutes{
		Source: source,
		Routes: make([]RouteInfo, 0, len(dists)),
	}
	for _, dist := range dists {
		routes, err := getAllRoutes(source, dist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result.Routes = append(result.Routes, routes.Routes...)
	}
	fmt.Fprint(w, result)
	json.NewEncoder(w).Encode(result)
}

func handleRequest() {
	http.HandleFunc("/routes", getRoutes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequest()
}
