package main

import (
	"database/sql"
	"fmt"

	_ "github.com/Go-SQL-Driver/MYSQL"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mia?charset=utf8")
	if err != nil {
		fmt.Println("sql open error")
		fmt.Println(err)
		return
	}

	if err := db.Ping(); err != nil {
		fmt.Println("%s error ping database: %s", err.Error())
		return
	}

	rows, err := db.Query("SELECT id,type as typea from wx_ad_feedback")
	//普通demo
	for rows.Next() {
		println(" into row")
		var id int
		var typea string
		defer rows.Close()
		//rows.Columns()
		err = rows.Scan(&id, &typea)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
		fmt.Println(typea)
	}
}

//db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
