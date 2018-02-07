package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123@/TESTDB?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
	fmt.Print("init ok")
}

func main() {
	http.HandleFunc("/user", userInfo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listenandserve:", err)
	}
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	query()
	fmt.Fprint(w, "查询完成")
}

func insert() {
	stmt, err := db.Prepare("insert user (user_name,user_age,user_sex) values (?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func query() {
	rows, err := db.Query("select * from user")
	checkErr(err)

	//获取列表字段
	columns, _ := rows.Columns()
	//构造字典类型
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//保存数据到字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				//得到的数据存入map
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func update() {
	stmt, err := db.Prepare("update user set user_age = ?, user_sex=? where user_id=?")
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func remove() {
	stmt, err := db.Prepare("delete from user whers user_id=?")
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)

	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
