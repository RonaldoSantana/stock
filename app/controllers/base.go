package controllers

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

// errorResponse type
type errorResponse struct {
	Error string `json:"Error"`
}

// response type
type response struct {
	Text     string      `json:"text"`
	Data     interface{} `json:"data"`
	Redirect string      `json:"redirect"`
}

// status type
type status struct {
	App     string `json:"app"`
	Version string `json:"version"`
}

// jsonStatus type
type jsonStatus struct {
	Code     int    `json:"code"`
	Text     string `json:"text"`
	Redirect string `json:"redirect"`
}

// sessionTokenResponse type
type sessionTokenResponse struct {
	Token string `json:"sessionToken"`
}

// baseController type
type baseController struct {
	Env         map[string]interface{}
	Session     *sessions.Session
	DB          *sql.DB
	EmailClient *helper.EmailClient
}

// Index func
func (bc *baseController) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var status status
	status.App = bc.Env["App"].(string)
	status.Version = bc.Env["Version"].(string)
	httpWrite(writer, request, status, http.StatusOK)
}

// writer
func httpWrite(writer http.ResponseWriter, request *http.Request, body interface{}, code int) {
	err := bc.Session.Save(request, writer)
	if err != nil {
		log.Println(err)
		HTTPError(writer, http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Expose-Headers", "X-API-Version")
	writer.Header().Set("X-API-Version", bc.Env["APIVersion"].(string))

	if _, exists := bc.Env["Location"]; exists {
		writer.Header().Set("Location", bc.Env["Location"].(string))
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

// HTTPError func
func HTTPError(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(jsonStatus{Code: code, Text: http.StatusText(code)}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, http.StatusText(http.StatusInternalServerError))
		return
	}
}

// HTTPWriteJSON func
func HTTPWriteJSON(writer http.ResponseWriter, status int, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(response)
}

// HTTPWriteError func
func HTTPWriteError(writer http.ResponseWriter, status int, response interface{}) error {
	return HTTPWriteJSON(writer, status, errorResponse{Error: response.(string)})
}

// HTTPNotFound func
func HTTPNotFound(wrt http.ResponseWriter, req *http.Request) {
	wrt.WriteHeader(http.StatusNotFound)
	httpWrite(wrt, req, "errors/404", http.StatusNotFound)
}
