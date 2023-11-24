package config

import (
	"crud_go_native/helpers"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_products?parseTime=true")
	helpers.FuncError(err, "error at ConnectDB")

	DB = db
}
