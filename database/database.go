package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	// "github.com/ChalanthornA/katradebutgo/entity"
)

var DB *sql.DB;


func Connect() error{
	var err error;
	dsn := "host=localhost user=Palm password=Palm_p2001 dbname=katrade port=5434 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = sql.Open("postgres", dsn)
	fmt.Println("Hi DB")

	if err != nil {
		return err;
	}
	fmt.Println("connect to db")
	return nil;
}
