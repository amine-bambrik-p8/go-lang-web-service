package routes

import (
	"encoding/json"
	"fmt"
	routesModel "go-lang-web-service/models/routes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

const BaseURL = "http://router.project-osrm.org/route/v1/driving/"

func GetAllRoutes(source string, dist string) (routes *routesModel.AllRoutes, err error) {
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
