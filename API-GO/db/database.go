package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {

	cnx := "****:****!@tcp(localhost:3306)/****"

	var err error
	DB, err = sql.Open("mysql", cnx)
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexión Exitosa")

}

func CloseConnection() {
	if DB != nil {
		DB.Close()
	}
}
