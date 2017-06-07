package main

import (
	"fmt"
	"github.com/rmsj/stock/db/models"
	"database/sql"
	"github.com/vattle/sqlboiler/queries/qm"
)

func main() {
	
	// Open handle to database like normal
	db, err := sql.Open("mysql", "stock_dba:secret@tcp(192.168.10.10:3306)/stock?charset=utf8")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	
	// SELECT COUNT(*) FROM pilots;
	count, err := models.Countries(db, qm.Limit(5)).Count()
	if err != nil {
		fmt.Errorf("error: %v", db)
	}
	fmt.Println(count)
}
