package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// CREATE TABLE `t1` (
//  `f1` int(11) NOT NULL,
//  `f2` bigint(20) DEFAULT NULL,
//  `f3` double DEFAULT NULL,
//  `f4` datetime DEFAULT NULL,
//  `f5` timestamp NULL DEFAULT NULL,
//  `f6` varchar(255) DEFAULT NULL,
//  `f7` longtext,
//  `f8` longblob,
//  `f9` bit(1) DEFAULT NULL,
//  PRIMARY KEY (`f1`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

func main() {
	// Open database connection
	db, err := sql.Open("mysql", "root:XXXXXX@tcp(192.168.5.19:3306)/Twdata")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i, col := range columns {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if (ok) {
				v = string(b)
			} else {
				v = val
			}

			fmt.Println(col, v)
		}
	}
}