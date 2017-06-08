package app

import (
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"database/sql"
	
	"github.com/gorilla/sessions"
	"gopkg.in/yaml.v2"
	"github.com/rmsj/stock/app/helpers"
)


// Application is the main app
type Application struct {
	Env       	map[string]interface{}
	Port      	string
	Config    	string
	Store   	*sessions.CookieStore
	DB			*sql.DB
	EmailClient *helpers.EmailClient
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

	app.setCookieStore()
	app.setDefaultGlobals()
	app.setMiddleware()

	log.Println("Starting server on port", app.Port)
	log.Fatal(http.ListenAndServe(":"+app.Port, &RouterCORS{app.NewRouter()}))
}
