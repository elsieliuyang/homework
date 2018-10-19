package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	uhostUserName = "uhost"
	uhostPassword = "uhost"
	uhostIP       = "172.28.217.246"
	uhostDBName   = "u_host"
)

func openDB(username string, password string, ip string, dbname string) (*sql.DB, error) {
	if len(username) == 0 {
		log.Fatalf("username is empty")
	}
	if len(password) == 0 {
		log.Fatalf("password is empty")
	}
	if len(ip) == 0 {
		log.Fatalf("ip is empty")
	}
	if len(dbname) == 0 {
		log.Fatalf("dbname is empty")
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, ip, dbname)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("sql.open failed: %v", err)
	}
	return db, err
}

func checkuhostinstancedata() {

	db, err := openDB(uhostUserName, uhostPassword, uhostIP, uhostDBName)
	if err != nil {
		log.Fatalf("openDB uaccount failed: %v", err)
	}
	defer db.Close()

	sql := "select create_time from uhostinstance where to_days(now()) - to_days(from_unixtime(create_time)) <= 1"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalf("db.Query failed: %v", err)
	}

	if rows == nil {
		fmt.Println("error: uhostinstance update fail")
	}

	var newuhostinstance []string
	for rows.Next() {
		var createTime string
		err := rows.Scan(&createTime)
		if err != nil {
			log.Println(err)
		}
		newuhostinstance = append(newuhostinstance, createTime)
	}

	if rows != nil {
		fmt.Println("uhostinstance update succeed:", len(newuhostinstance))
	}
	defer rows.Close()
}

func main() {
	checkuhostinstancedata()
}
