package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestJSON(url string, response interface{}) (err error) {
	res, err := http.Get(url)
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
