package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

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

var (
	dbTUba  *sql.DB
	dbTWant *sql.DB
)

func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != sql.ErrTxDone && err != nil {
		logger.Panicf("clearTransaction -> Error: %s \n", err)
	}
}

// 初始化數據庫
func initMysql() bool {
	conn1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Mysqls.Tuba.Username, config.Mysqls.Tuba.Password, config.Mysqls.Tuba.Host, config.Mysqls.Tuba.Port, config.Mysqls.Tuba.Database)
	db1, err1 := sql.Open("mysql", conn1)
	if err1 != nil {
		logger.Panicf("Fatal error2 initLogger: %s \n", err1)
		return false
	}
	dbTUba = db1
	conn2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Mysqls.TWant.Username, config.Mysqls.TWant.Password, config.Mysqls.TWant.Host, config.Mysqls.TWant.Port, config.Mysqls.TWant.Database)
	db2, err2 := sql.Open("mysql", conn2)
	if err2 != nil {
		logger.Panicf("Fatal error2 initLogger: %s \n", err2)
		return false
	}
	dbTWant = db2
	return true
}

// 反初始化
func uninitialized() {
	if dbTUba != nil {
		_ = dbTUba.Close()
	}
	if dbTWant != nil {
		_ = dbTWant.Close()
	}
}

// execute sql
func dbExecuteByTrans(tx *sql.Tx, sql string, values ...interface{}) (int64, error) {
	stmtIns, err1 := tx.Prepare(sql) // ? = placeholder
	if err1 != nil {
		return -1, fmt.Errorf("dbExecuteByTrans -> Error1: %s - %s - %v", err1, sql, values)
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	result, err2 := stmtIns.Exec(values...)
	if err2 != nil {
		return -1, fmt.Errorf("dbExecuteByTrans -> Error2: %s - %s - %v", err2, sql, values)
	}
	return result.RowsAffected()
}

func dbExecute(db *sql.DB, sql string, rowStruct interface{}) (int64, error) {
	stmtIns, err1 := db.Prepare(sql) // ? = placeholder
	if err1 != nil {
		logger.Printf("dbExecute -> Error1: %s - %s - %v", err1, sql, rowStruct)
		return -1, err1
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	s := reflect.ValueOf(rowStruct).Elem()
	row := make([]interface{}, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		row[i] = s.Field(i).Addr().Interface()
	}
	result, err2 := stmtIns.Exec(row...)
	if err2 != nil {
		logger.Printf("dbExecute -> Error2: %s - %s - %v", err2, sql, rowStruct)
		return -1, err2
	}
	return result.RowsAffected()
}

func dbExecutes(db *sql.DB, sql string, values ...interface{}) (int64, error) {
	stmtIns, err1 := db.Prepare(sql) // ? = placeholder
	if err1 != nil {
		logger.Printf("dbExecute -> Error1: %s - %s - %v", err1, sql, values)
		return -1, err1
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	result, err2 := stmtIns.Exec(values...)
	if err2 != nil {
		logger.Printf("dbExecute -> Error2: %s - %s - %v", err2, sql, values)
		return -1, err2
	}
	return result.RowsAffected()
}

func dbGetColumnTypes(db *sql.DB, name string) ([]sql.ColumnType, error) {
	sqlSelect := fmt.Sprintf("SELECT * FROM %s WHERE 1=2", name)
	rows, err1 := db.Query(sqlSelect)
	if err1 != nil {
		logger.Printf("dbGetColumnTypes -> Error1: %s \n", err1)
		return nil, err1
	}
	if ctypes, err2 := rows.ColumnTypes(); err2 != nil {
		logger.Printf("dbGetColumnTypes -> Error2: %s \n", err2)
		return nil, err2
	} else {
		columnTypes := make([]sql.ColumnType, 0, len(ctypes))
		for i := 0; i < len(ctypes); i++ {
			columnTypes = append(columnTypes, *ctypes[i])
		}
		return columnTypes, nil
	}
}

func dbQueryRow(db *sql.DB, query string, rowStruct interface{}) error {
	rows, err1 := db.Query(query)
	if err1 != nil {
		logger.Printf("dbQueryRow -> Error1: %s \n", err1)
		return err1
	}
	defer rows.Close()
	s := reflect.ValueOf(rowStruct).Elem()
	row := make([]interface{}, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		row[i] = s.Field(i).Addr().Interface()
	}
	for rows.Next() {
		if err2 := rows.Scan(row...); err2 != nil {
			logger.Printf("dbQueryRow -> Error2: %s \n", err2)
			return err2
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
	if err3 := rows.Err(); err3 != nil {
		logger.Printf("dbQueryRow -> Error3: %s \n", err3)
		return err3
	}
	return nil
}

func dbQueryRows(db *sql.DB, query string, fCreateRow func() interface{}) error {
	rows, err1 := db.Query(query)
	if err1 != nil {
		logger.Printf("dbQueryRow -> Error1: %s \n", err1)
		return err1
	}
	defer rows.Close()
	for rows.Next() {
		v := fCreateRow()
		s := reflect.ValueOf(v).Elem()
		row := make([]interface{}, s.NumField())
		for i := 0; i < s.NumField(); i++ {
			row[i] = s.Field(i).Addr().Interface()
		}
		if err2 := rows.Scan(row...); err2 != nil {
			logger.Printf("dbQueryRow -> Error2: %s \n", err2)
			return err2
		}
	}
	if err3 := rows.Err(); err3 != nil {
		logger.Printf("dbQueryRow -> Error3: %s \n", err3)
		return err3
	}
	return nil
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

func helloDbQueryRow() {
	db, err := sql.Open("mysql", "root:XXXXXX(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	t1 := T1{}
	dtNow1 := time.Now()
	dbQueryRow(db, "SELECT * FROM t1", &t1)
	fmt.Println("--- -------------------- ---", tmDiff2Now(dtNow1))
	fmt.Println(t1)
}

func helloDbQueryRows() {
	var rs []T1
	createT1 := func() interface{} {
		rs = append(rs, T1{})
		return &rs[len(rs)-1]
	}
	db, err := sql.Open("mysql", "root:XXXXXX(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	dtNow1 := time.Now()
	dbQueryRows(db, "SELECT * FROM t1", createT1)
	fmt.Println("--- -------------------- ---", tmDiff2Now(dtNow1))
	fmt.Println(len(rs))
}

func helloDbExecute() {
	db, err := sql.Open("mysql", "root:XXXXXX(192.168.5.19:3306)/Twdata?charset=utf8&parseTime=true")
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
