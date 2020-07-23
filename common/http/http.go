package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/amine-bambrik-p8/go-lang-web-service/models"
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

// Takes data and status code and sends it as JSON
func SendJSON(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	var buf bytes.Buffer

	if err, ok := data.(error); ok {
		data = map[string]interface{}{
			"error": err.Error(),
		}
	}
	if obj, ok := data.(models.Model); ok {
		data = obj.GetViewModel()
	}
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("Respond:", err)
	}
}
