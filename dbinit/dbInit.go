package dbinit

import (
	"database/sql"
	"fmt"
)

func DbInit() *sql.DB {
var db *sql.DB
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?parseTime=true&&charset=utf8")
	if err!=nil{
	fmt.Println(err,"db connect fail")
	}
	return db
}


