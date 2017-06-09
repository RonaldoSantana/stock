package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/rmsj/stock/app/helper"
)

// API type
type API struct {
	Env         map[string]interface{}
	Session     *sessions.Session
	DB          *sql.DB
	EmailClient *helper.EmailClient
}

// ErrorResponse type
type ErrorResponse struct {
	Error string `json:"Error"`
}

// Response type
type Response struct {
	Text     string      `json:"text"`
	Data     interface{} `json:"data"`
	Redirect string      `json:"redirect"`
}

// Status type
type Status struct {
	App     string `json:"app"`
	Version string `json:"version"`
}

// jsonStatus type
type jsonStatus struct {
	Code     int    `json:"code"`
	Text     string `json:"text"`
	Redirect string `json:"redirect"`
}

// Index func
func (api *API) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var status Status
	status.App = api.Env["App"].(string)
	status.Version = api.Env["Version"].(string)
	api.writer(writer, request, status, http.StatusOK)
}

func (api *API) writer(writer http.ResponseWriter, request *http.Request, body interface{}, code int) {
	err := api.Session.Save(request, writer)
	if err != nil {
		log.Println(err)
		api.Error(writer, http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Expose-Headers", "X-API-Version")
	writer.Header().Set("X-API-Version", api.Env["APIVersion"].(string))

	if _, exists := api.Env["Location"]; exists {
		writer.Header().Set("Location", api.Env["Location"].(string))
	}

	writer.WriteHeader(code)
	b, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, http.StatusText(http.StatusInternalServerError))
		return
	}

	_, err = writer.Write(b)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (api *API) Error(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(jsonStatus{Code: code, Text: http.StatusText(code)}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func (api *API) WriteJSON(writer http.ResponseWriter, status int, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(response)
}

func (api *API) WriteError(writer http.ResponseWriter, status int, response interface{}) error {
	return api.WriteJSON(writer, status, ErrorResponse{Error: response.(string)})
}

// NotFound func
func (api *API) NotFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	api.writer(writer, request, "errors/404", http.StatusNotFound)
}
