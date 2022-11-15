package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	db, _ := sql.Open("mysql", "root:admin.1234@tcp(127.0.0.1:3306)/gotest")
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	defer db.Close()
}


//
//func ConnectMysql() *sql.DB {
//	mysqlUrl := "root:admin.1234@tcp(127.0.0.1:3306)/gotest"
//	db, _ := sql.Open("mysql", mysqlUrl)
//	defer db.Close()
//	err := db.Ping()
//	if err != nil{
//		fmt.Println("mysql连接失败，错误日志为：", err.Error())
//		return nil
//	}else{
//		// Do something here
//	}
//	return db
//}
