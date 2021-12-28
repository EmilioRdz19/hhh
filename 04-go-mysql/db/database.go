package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database
const url = "root:1234@tcp(localhost:3306)/libreria"

//Guarda la conexion
var db *sql.DB

//Realiza la conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection
}

func Close() {
	db.Close()
}

//Verifica
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
