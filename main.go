package main

import (
	"fmt"
	"database/sql"
	
	"github.com/rmsj/stock/db/models"
	_ "github.com/go-sql-driver/mysql"
)

//go:generate sqlboiler --basedir "./db/" --tinyint-as-bool mysql

func main() {
	
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
