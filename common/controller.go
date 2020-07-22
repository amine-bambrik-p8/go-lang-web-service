package common

import (
	"bytes"
	"encoding/json"
	"go-lang-web-service/models"
	"io"
	"log"
	"net/http"
)

type Controller struct {
}

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
