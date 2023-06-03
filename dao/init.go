package dao

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// DB接続のため環境変数を取得(本番環境用）
	/*
		mysqlUser := os.Getenv("MYSQL_USER")
		mysqlPwd := os.Getenv("MYSQL_PASSWORD")
		mysqlHost := os.Getenv("MYSQL_HOST")
		mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	*/

	// DB接続のため環境変数を取得（開発環境用）
	err := godotenv.Load("../database/.env")
	if err != nil {
		fmt.Printf("読み込みできませんでした：%v", err)
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	fmt.Println("user:", mysqlUser)
	fmt.Println("password:", mysqlPwd)
	fmt.Println("database:", mysqlDatabase)

	//本番環境用
	//connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	//開発環境用
	connStr := fmt.Sprintf("%s:%s@(localhost:3308)/%s", mysqlUser, mysqlPwd, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	//
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}
