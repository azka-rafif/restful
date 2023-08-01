package response

import (
	"encoding/json"
	"go-mini/internal/log"
	"go-mini/internal/response/failure"
	"net/http"
)

type Base struct {
	Data    *interface{} `json:"data,omitempty"`
	Error   *string      `json:"error,omitempty"`
	Message *string      `json:"message,omitempty"`
}

func WithJSON(w http.ResponseWriter, code int, jsonPayload interface{}) {
	respond(w, code, Base{Data: &jsonPayload})
}

func NoContent(w http.ResponseWriter) {
	respond(w, http.StatusNoContent, nil)
}

func WithError(w http.ResponseWriter, err error) {
	code := failure.GetCode(err)
	errMsg := err.Error()
	respond(w, code, Base{Error: &errMsg})
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	logger := log.GetLogger()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		logger.Error(err)
	}
}
