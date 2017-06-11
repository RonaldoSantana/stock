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

var config map[string]interface{}
var session *sessions.Session
var db *sql.DB
var email *helper.EmailClient

// Initialize controllers
func Initialize(envMap map[string]interface{}, appSession *sessions.Session, dataBase *sql.DB, mailClient *helper.EmailClient) {
	config = envMap
	session = appSession
	db = dataBase
	email = mailClient
}

// Index func
func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var status status
	status.App = config["App"].(string)
	status.Version = config["Version"].(string)
	write(writer, request, status, http.StatusOK)
}

// writer
func write(writer http.ResponseWriter, request *http.Request, body interface{}, code int) {
	err := session.Save(request, writer)
	if err != nil {
		log.Println(err)
		Error(writer, http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Expose-Headers", "X-API-Version")
	writer.Header().Set("X-API-Version", config["APIVersion"].(string))

	if _, exists := config["Location"]; exists {
		writer.Header().Set("Location", config["Location"].(string))
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

// Error func
func Error(writer http.ResponseWriter, code int) {
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(jsonStatus{Code: code, Text: http.StatusText(code)}); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, http.StatusText(http.StatusInternalServerError))
		return
	}
}

// WriteJSON func
func WriteJSON(writer http.ResponseWriter, status int, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(response)
}

// WriteError func
func WriteError(writer http.ResponseWriter, status int, response interface{}) error {
	return WriteJSON(writer, status, errorResponse{Error: response.(string)})
}

// HTTPNotFound func
func NotFound(wrt http.ResponseWriter, req *http.Request) {
	wrt.WriteHeader(http.StatusNotFound)
	write(wrt, req, "errors/404", http.StatusNotFound)
}
