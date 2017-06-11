package app

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gopkg.in/yaml.v2"
)

// Application is the main app
type Application struct {
	Env    map[string]interface{}
	Port   string
	Config string
	Routes []Route
	Store  *sessions.CookieStore
	DB     *sql.DB
}

func (app *Application) setDefaultGlobals() {

	app.Env["APP_ENV"] = "dev"
	if os.Getenv("APP_ENV") != "" {
		app.Env["APP_ENV"] = os.Getenv("APP_ENV")
	}
	if app.Env["APP_ENV"].(string) == "prod" {
		app.Env["Host"] = app.Env["HostProd"].(string)
	}
	if app.Env["APP_ENV"].(string) == "test" {
		app.Env["Host"] = app.Env["HostTest"].(string)
	}

	app.Env["Version"] = app.GetVersion()
}

func (app *Application) setCookieStore() {
	store := sessions.NewCookieStore([]byte(app.Env["Secret"].(string)))
	app.Store = store
}

// loads application config file
func (app *Application) loadConfig() error {
	b, err := ioutil.ReadFile(app.Config)
	if err != nil {
		return err
	}
	return yaml.Unmarshal([]byte(b), &app.Env)
}

// Run application - creating DB session, etc.
func (app *Application) Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := app.loadConfig(); err != nil {
		log.Fatal(err)
	}
	if err := app.connectDB(); err != nil {
		log.Fatal(err)
	}
	defer app.DB.Close()

	app.setCookieStore()
	app.setDefaultGlobals()
	app.setMiddleware()

	log.Println("Starting server on port", app.Port)
	log.Fatal(http.ListenAndServe(":"+app.Port, &RouterCORS{app.NewRouter()}))
}
