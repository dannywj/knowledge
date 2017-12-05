package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/Go-SQL-Driver/MYSQL"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mia?charset=utf8")
	if err := db.Ping(); err != nil {
		fmt.Println("%s error ping database: %s", err.Error())
		return
	}
}

func main() {
	println("======begin mysql DB Demo======")
	baseGet()
	mapGet()
}

func baseGet() {
	rows, err := db.Query("SELECT id,username from user_info")
	//普通demo
	println("--base method--")
	for rows.Next() {
		var id int
		var username string
		defer rows.Close()
		//fmt.Println(rows.Columns())
		err = rows.Scan(&id, &username)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("id:" + strconv.Itoa(id))
		fmt.Println("username:" + username)
	}
}

func mapGet() {
	rows, _ := db.Query("SELECT * from user_info")
	println("--map method--")
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println("map struct:")
		fmt.Println(record)
		fmt.Println("map data:")
		fmt.Println("username:" + record["username"])
	}
}

/*
======begin mysql DB Demo======
--base method--
id:1
username:wangjue@mia.com
id:2
username:test@111.com
--map method--
map struct:
map[id:1 username:wangjue@mia.com password:e10adc3949ba59abbe56e057f20f883e create_time:2017-07-20 10:46:55 type:1 status:0]
map data:
username:wangjue@mia.com
map struct:
map[id:2 username:test@111.com password:e10adc3949ba59abbe56e057f20f883e create_time:2017-07-24 11:01:17 type:0 status:0]
map data:
username:test@111.com
*/
