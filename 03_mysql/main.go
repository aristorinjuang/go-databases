package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	ID      int
	Title   string
	Content string
}

var articles []article

func main() {
	db, err := sql.Open("mysql", "root:root@/go_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	rows, err := db.Query(`SELECT * FROM articles;`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id      int
			title   string
			content string
		)
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			log.Println(err)
		}
		articles = append(articles, article{id, title, content})
	}

	fmt.Println(articles)
}
