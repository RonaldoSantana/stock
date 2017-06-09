package main

import (
	/*"os"
	"fmt"
	"flag"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
	"github.com/rmsj/stock/db/models"
	"github.com/rmsj/stock/app"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"*/
)



import (
	"fmt"
	"flag"
	"os"
	
	"github.com/rmsj/stock/app"
)

//go:generate sqlboiler -o "./db/models/" --tinyint-as-bool mysql
/*

func fake_main() {
	
	// Open handle to database like normal
	conn, err := sql.Open("mysql", "stock_dba:secret@tcp(192.168.10.10:3306)/stock?charset=utf8")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	
	// SELECT COUNT(*) FROM pilots;
	count, err := models.Countries(conn).Count()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Println(count)
}
*/

// App is the main application object
var App app.Application

func init() {
	var rev, version bool
	flag.StringVar(&App.Port, "port", "8080", "Port to listen")
	flag.StringVar(&App.Config, "config", "config.yml", "Config yaml file")
	flag.BoolVar(&rev, "rev", false, "Set revision")
	flag.BoolVar(&version, "version", false, "Display version")
	flag.Parse()
	
	if rev != false {
		App.SetRevision()
		os.Exit(0)
	}
	
	if version != false {
		ver := App.GetVersion()
		fmt.Println(ver)
		os.Exit(0)
	}
}

func main() {
	App.Run()
}