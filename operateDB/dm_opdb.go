package main

import (
	"database/sql"
	"fmt"

	_ "github.com\go-sql-driver\mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	stmt, err := db.Prepare("INSERT userinfo SET user.name=?,departname=?,created=?")

	res, err := stmt.Exec("Kris Shin", "Development", "2020-5-25")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("Kris Shin very handsome", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)
	}

	stmt, err = db.Prepare("DELETE FROM userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
