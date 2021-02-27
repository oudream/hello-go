package main

import (
	"crypto"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strconv"
	"time"
)

func helloDb1() {
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1")
	//db, err := sql.Open("mysql", "user:password@/database")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO t1 VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT f2 FROM t1 WHERE f3 > ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(strconv.Itoa(i), i, i * i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d\n", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d\n", squareNum)
}

func helloSelect1() {
	// Open database connection
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		log.Panicf("clearTransaction -> Error: %s \n", err)
	}
}


func helloMysql1() {
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer clearTransaction(tx)

	rs, err := tx.Exec("UPDATE user SET gold=50 WHERE real_name='vanyarpy'")
	if err != nil {
		log.Fatalln(err)
	}
	rowAffected, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rowAffected)

	rs, err = tx.Exec("UPDATE user SET gold=150 WHERE real_name='noldorpy'")
	if err != nil {
		log.Fatalln(err)
	}
	rowAffected, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rowAffected)

	if err := tx.Commit(); err != nil {
		// tx.Rollback() 此时处理错误，会忽略doSomthing的异常
		log.Fatalln(err)
	}

	//"{\"code\":200,\"message\":\"OK\",\"data\":\"\"}"

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO t1 VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT f2 FROM t1 WHERE f3 > ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(strconv.Itoa(i), i, i*i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d\n", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d\n", squareNum)
}

type T1 struct {
	f1 sql.NullInt32
	f2 sql.NullFloat64
	f3 sql.NullString
	f4 sql.NullString
	f5 sql.NullString
}

type T2 struct {
	f1 int
	f2 float64
	f3 *time.Time
	f4 bool
	f5 string
}

type T3 struct {
	f1 *int
	f2 *float64
	f3 *time.Time
	f4 []byte
	f5 *string
}

type T4 struct {
	f1 *int
	f2 *float64
	f3 *time.Time
	f4 []byte
	f5 *string
}

var (
	layout     = flag.String("layout", "", "")
)

// unsupported Scan, storing driver.Value type []uint8 into type *time.Time
// https://blog.csdn.net/han0373/article/details/81698713
// ?parseTime=true
func helloSelect2() {
	// Open database connection
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1?parseTime=true")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var (
			f1 *int
			f2 *float64
			f3 *time.Time
			f4 *[]byte // bool
			f5 *string
		)
		_ = rows.Scan(&f1, &f2, &f3, &f4, &f5)
		var t2 T2
		if f1 != nil {
			t2.f1 = *f1
		} else {
			t2.f1 = 0
		}
		if f2 != nil {
			t2.f2 = *f2
		} else {
			t2.f2 = 0
		}
		if f3 != nil {
			t2.f3 = f3
		} else {
			t2.f3 = &time.Time{}
		}
		if f4 != nil {
			t2.f4 = (*f4)[0] > 0
		} else {
			t2.f4 = false
		}
		if f5 != nil {
			t2.f5 = *f5
		} else {
			t2.f5 = ""
		}
		fmt.Println(t2)
		//var data []byte = []byte(t2.f4)
		//layout := string(data[:])
		//tm,_ := time.Parse(time.RFC3339, t2.f4)
		//tm,_ := time.Parse("2006-01-02 15:04:05", t2.f5)
		//tm,_ := time.Parse(layout, t2.f4)
		//fmt.Println(tm)
	}
}

func helloSelect3() {
	// Open database connection
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1?parseTime=true")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var t3 T3
		_ = rows.Scan(&t3.f1, &t3.f2, &t3.f3, &t3.f4, &t3.f5)
		fmt.Println(t3)
		//var data []byte = []byte(t2.f4)
		//layout := string(data[:])
		//tm,_ := time.Parse(time.RFC3339, t2.f4)
		//tm,_ := time.Parse("2006-01-02 15:04:05", t2.f5)
		//tm,_ := time.Parse(layout, t2.f4)
		//fmt.Println(tm)
	}
}

func helloSelect4() {
	// Open database connection
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1?parseTime=true")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var t3 T4
		_ = rows.Scan(&t3.f1, &t3.f2, &t3.f3, &t3.f4, &t3.f5)
		fmt.Println(t3)
		//var data []byte = []byte(t2.f4)
		//layout := string(data[:])
		//tm,_ := time.Parse(time.RFC3339, t2.f4)
		//tm,_ := time.Parse("2006-01-02 15:04:05", t2.f5)
		//tm,_ := time.Parse(layout, t2.f4)
		//fmt.Println(tm)
	}
}

func dbGetColumnTypes1(name string, db *sql.DB) []*sql.ColumnType {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE 1=2", name)
	rows, err1 := db.Query(sql)
	if err1 != nil {
		panic(err1.Error()) // proper error handling instead of panic in your app
	}
	if columnTypes, err2 := rows.ColumnTypes(); err2 != nil {
		panic(err2.Error()) // proper error handling instead of panic in your app
	} else {
		return columnTypes
	}
}

func dbGetColumnTypes2(name string, db *sql.DB) []*sql.ColumnType {
	sql := fmt.Sprintf("show full columns from %s", name)
	rows, err1 := db.Query(sql)
	if err1 != nil {
		panic(err1.Error()) // proper error handling instead of panic in your app
	}
	if columnTypes, err2 := rows.ColumnTypes(); err2 != nil {
		panic(err2.Error()) // proper error handling instead of panic in your app
	} else {
		return columnTypes
	}
}

func Hash(objs ...interface{}) []byte {
	digester := crypto.MD5.New()
	for _, ob := range objs {
		fmt.Fprint(digester, reflect.TypeOf(ob))
		fmt.Fprint(digester, ob)
	}
	return digester.Sum(nil)
}

func helloSelect5() {
	// Open database connection
	db, err := sql.Open("mysql", "root:Aa.XXXXXX@tcp(192.168.5.110:3306)/db1?parseTime=true")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columnTypes, _ := rows.ColumnTypes()
	for i := 0; i < len(columnTypes); i++ {
		fmt.Println(columnTypes[i])
	}

	vs := make(map[string]map[string]interface{})
	for rows.Next() {
		row := make(map[string]interface{})
		for i := 0; i < len(columnTypes); i++ {
			var v interface{}
			_ = rows.Scan(&v)
			row[columnTypes[i].Name()] = v
		}
		//vs[]
		//_ = rows.Scan(&t3.f1, &t3.f2, &t3.f3, &t3.f4, &t3.f5)
		//fmt.Println(t3)
		//var data []byte = []byte(t2.f4)
		//layout := string(data[:])
		//tm,_ := time.Parse(time.RFC3339, t2.f4)
		//tm,_ := time.Parse("2006-01-02 15:04:05", t2.f5)
		//tm,_ := time.Parse(layout, t2.f4)
		//fmt.Println(tm)
	}
	fmt.Println(vs)
}

func helloSelect6() {
	// Open database connection
	db, err := sql.Open("mysql", "rds_user1:Aa@XXXXXX@tcp(rm-j6cl7dib708j6i6i5to.mysql.rds.aliyuncs.com:3306)/db1?parseTime=true")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	now := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		dbGetColumnTypes1("t1", db)
	}
	fmt.Println(time.Now().UnixNano() - now)
	now = time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		dbGetColumnTypes2("t1", db)
	}
	fmt.Println(time.Now().UnixNano() - now)
}

func main() {
	helloSelect5()
}
