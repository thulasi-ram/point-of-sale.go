package web

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data   interface{}   `json:"data,omitempty"`
	Errors []interface{} `json:"errors,omitempty"`
	Status int           `json:"-"`
}

type Endpoint func(app *App, req *http.Request) *APIResponse
type EndpointHandler func(e Endpoint) http.HandlerFunc

func NewAPIResponse(data interface{}, status int) *APIResponse {
	return &APIResponse{Data: data, Status: status}
}

func NewErrorAPIResponse(err error, status int) *APIResponse {
	return &APIResponse{Errors: []interface{}{err.Error()}, Status: status}
}

func writeAPIResponse(w http.ResponseWriter, resp *APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

func HandlerFunc(app *App) EndpointHandler {
	return func(e Endpoint) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			response := e(app, request)
			writeAPIResponse(writer, response)
		}
	}
}
