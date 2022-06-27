package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() {

	db, err := sql.Open("mysql", "root:@/blogGo")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connecting to database")
		return
	}
	fmt.Println("Connected to database")

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database")
		return
	} else {
		fmt.Println("Pinged database")
	}
}

// func connToString(info Connection) string {
// 	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", info.Host, info.Port, info.User, info.Password, info.DBName)
// }
