package app

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (app *Application) connectDB() error {

	db, err := sql.Open(app.Env["DBDriver"].(string), app.Env["DBDataSource"].(string))
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	// sets app DB
	app.DB = db

	return nil
}
