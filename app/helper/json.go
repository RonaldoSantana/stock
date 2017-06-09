package helper

import (
	"encoding/json"
	"io"
	"log"
)

// LoadFromJSON loads json from http request body
func LoadFromJSON(body io.Reader, i interface{}) (err error) {
	if err = json.NewDecoder(body).Decode(i); err != nil {
		log.Printf("Could not decode JSON: %v", err)
	}
	return
}
