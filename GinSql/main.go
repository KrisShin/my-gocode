package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := "root:qweqwe@tcp(localhost:3306)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 执行sql语句
	//_, err = db.Exec("create table person(id int auto_increment primary key," +
	//	"name varchar(12) not null," +
	//	"age int default 1);")
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//} else {
	//	fmt.Println("create table person success")
	//}

	// insert data
	//_, err = db.Exec("insert into person(name, age) "+
	//	"values(?,?);", "Kris", 22)
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//} else {
	//	fmt.Println("insert data success")
	//}

	// search data
	rows, err := db.Query("select id, name, age from person;")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

scan:
	if rows.Next() {
		person := new(Person)
		err = rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}
}

type Person struct {
	Id   int
	Name string
	Age  int
}
