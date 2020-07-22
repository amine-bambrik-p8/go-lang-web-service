package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPClient var for mocking requests
var (
	Client HTTPClient
)

// Create
func init() {
	Client = &http.Client{}
}

// Get sends a get request to the URL with the specified headers
// Returns *http.Response or error in case of failure
func Get(url string, headers http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		request.Header = headers
	}
	return Client.Do(request)

}

// Post sends a post request to the URL with the body & headers
// Returns *http.Response or error in case of failure
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}

// GetRequestJSON sends a post request and returns a response in the form of JSON
// Retruns an error in case of failure
func GetRequestJSON(url string, response interface{}) (err error) {
	res, err := Get(url, nil)
	if err != nil {
		log.Print("Failed to request routes")
		return
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Print("Failed to parse request body")
		return
	}

	if err = json.Unmarshal(bytes, &response); err != nil {
		log.Print("Error parsing json", err)
		return
	}
	return
}
