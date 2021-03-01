package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"time"
)

func queryRow(db *sql.DB, query string, rowStruct interface{}) {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 确定Scan函数的输入类型
	s := reflect.ValueOf(rowStruct).Elem()
	onerow := make([]interface{}, s.NumField())
	// 按顺序遍历结构体的每个元素，取其指针值
	for i := 0; i < s.NumField(); i++ {
		onerow[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		if err := rows.Scan(onerow...); err != nil {
			log.Fatal(err)
		}
		// 如果你要根据结果的数据类型做相应的处理，比如数据类型转换，比如把结果写到文件。。。
		//for i := 0; i < s.NumField(); i++ {
		//	switch s.Type().Field(i).Type.String() {
		//	case "time.Time":
		//		date := s.Field(i).Interface().(time.Time)
		//		fmt.Println(date.Format("2006-01-02"))
		//	default:
		//		fmt.Println("do nothing")
		//	}
		//}
		//fmt.Printf("parser succeed, result is: %v", s.Interface())
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func queryRows(db *sql.DB, query string, f func() interface{}) {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// 确定Scan函数的输入类型
		v := f()
		s := reflect.ValueOf(v).Elem()
		onerow := make([]interface{}, s.NumField())
		// 按顺序遍历结构体的每个元素，取其指针值
		for i := 0; i < s.NumField(); i++ {
			onerow[i] = s.Field(i).Addr().Interface()
		}
		if err := rows.Scan(onerow...); err != nil {
			log.Fatal(err)
		}
		// 如果你要根据结果的数据类型做相应的处理，比如数据类型转换，比如把结果写到文件。。。
		//for i := 0; i < s.NumField(); i++ {
		//	switch s.Type().Field(i).Type.String() {
		//	case "time.Time":
		//		date := s.Field(i).Interface().(time.Time)
		//		fmt.Println(date.Format("2006-01-02"))
		//	default:
		//		fmt.Println("do nothing")
		//	}
		//}
		//fmt.Printf("parser succeed, result is: %v", s.Interface())
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

// BitBool is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

func dbExecute(db *sql.DB, sql string, rowStruct interface{}) (int64, error) {
	stmtIns, err1 := db.Prepare(sql) // ? = placeholder
	if err1 != nil {
		return -1, fmt.Errorf("dbExecute -> Error1: %s - %s - %v", err1, sql, rowStruct)
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	s := reflect.ValueOf(rowStruct).Elem()
	row := make([]interface{}, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		row[i] = s.Field(i).Addr().Interface()
	}
	result, err2 := stmtIns.Exec(row...)
	if err2 != nil {
		return -1, fmt.Errorf("dbExecute -> Error2: %s - %s - %v", err2, sql, rowStruct)
	}
	return result.RowsAffected()
}

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
type T1 struct {
	F1 int
	F2 int64
	F3 float64
	F4 time.Time
	F5 time.Time
	F6 string
	F7 string
	F8 string
	F9 BitBool
}

type Goods struct {
	GoodsId int
	Price0 float64
	Price1 float64
	ImageName string
}

func main() {
	//hellosingle()
	//helloMulti()
	//helloGoods()
	helloDbExecute()
}

func tmDiff2Now(tm time.Time) int64 {
	return (time.Now().UnixNano() / 1e6) - (tm.UnixNano() / 1e6)
}

func helloSingle() {
	db, err := sql.Open("mysql", "root:XXXXXX@tcp(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	t1 := T1{}
	dtNow1 := time.Now()
	queryRow(db, "SELECT * FROM t1", &t1)
	fmt.Println("--- -------------------- ---", tmDiff2Now(dtNow1))
	fmt.Println(t1)
}

func helloMulti()  {
	var ts []T1
	createT1 := func() interface{} {
		ts = append(ts, T1{})
		return &ts[len(ts)-1]
	}
	db, err := sql.Open("mysql", "root:XXXXXX@tcp(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	queryRows(db, "SELECT * FROM t1", createT1)
	fmt.Println("--- -------------------- ---")
	fmt.Println(ts)
	//queryRows(db, "SELECT * FROM t1", createT1)
	//fmt.Println("--- -------------------- ---")
	//fmt.Println(ts)
}

func helloGoods()  {
	var rs []Goods
	createT1 := func() interface{} {
		rs = append(rs, Goods{})
		return &rs[len(rs)-1]
	}
	db, err := sql.Open("mysql", "root:XXXXXX@tcp(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	dtNow1 := time.Now()
	queryRows(db, "SELECT goods_id, app_price0, app_price1, image_name FROM goods", createT1)
	fmt.Println("--- -------------------- ---", tmDiff2Now(dtNow1))
	fmt.Println(len(rs))
}

func helloDbExecute() {
	db, err := sql.Open("mysql", "root:Ftofs7All#9@tcp(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	sqlInsert := "INSERT INTO `Twdata`.`t1`(`f1`, `f2`, `f3`, `f4`, `f5`, `f6`, `f7`, `f8`, `f9`) " +
		" VALUES " +
		" (?, ?, ?, ?, ?, ?, ?, ?, ?);"
	t1 := T1{}
	t1.F1 = 5
	t1.F6 = "a"
	t1.F7 = "b"
	t1.F8 = "c"
	dtNow1 := time.Now()
	r, _ := dbExecute(db, sqlInsert, &t1)
	fmt.Println("--- -------------------- ---", tmDiff2Now(dtNow1))
	fmt.Println(t1)
	fmt.Println(r)
}
