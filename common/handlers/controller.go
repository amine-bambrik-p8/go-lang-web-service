package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/amine-bambrik-p8/go-lang-web-service/models"
)

type Controller struct {
}

// Takes data and status code and sends it as JSON
func (c *Controller) SendJSON(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
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
