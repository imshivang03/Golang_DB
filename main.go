package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:imshivang03@tcp(127.0.0.1:3306)/employees")
	fmt.Println(db)
	if err != nil {
		panic(err.Error())
	}

	tx, _ := db.Begin()
	stmt, err := tx.Prepare("INSERT INTO employees(id,username) VALUES(?,?)")
	res, err := stmt.Exec(4, "Amit")
	res, err = stmt.Exec(56, "Lokesh")
  res, err := stmt.Exec(93, "Amritanshu")
	res, err = stmt.Exec(76, "Rahul")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	log.Println(res)

	rows, err := db.Query("select id, username from employees where id = ?", 4)
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
	)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	sqlStatementUpdt := `
    UPDATE employees
    SET username = "Gopal" WHERE id=33;`
	result, err := db.Exec(sqlStatementUpdt)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("rows updated: %v\n", count)

}
