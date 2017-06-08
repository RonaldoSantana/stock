package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Version type
type Version struct {
	Major     int       `json:"major"`
	Minor     int       `json:"minor"`
	Revision  int       `json:"revision"`
	Prefix    string    `json:"prefix"`
	Suffix    string    `json:"suffix"`
	API       string    `json:"api"`
	Timestamp time.Time `json:"timestamp"`
}

// GetVersion func
func (app *Application) GetVersion() (ver string) {
	var version Version
	data, err := ioutil.ReadFile("version.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &version)
	if err != nil {
		return
	}
	ver = fmt.Sprintf("%s%d.%d.%d%s", version.Prefix, version.Major, version.Minor, version.Revision, version.Suffix)
	return
}

// SetRevision func
func (app *Application) SetRevision() {
	var version Version
	data, err := ioutil.ReadFile("version.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &version)
	if err != nil {
		return
	}

	version.Revision = version.Revision + 1
	version.Timestamp = time.Now().UTC()

	jsonData, err := json.MarshalIndent(version, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile("version.json", jsonData, 0644)
	if err != nil {
		return
	}

	newVersion := app.GetVersion()
	log.Printf("Set version %s", newVersion)
	return
}
