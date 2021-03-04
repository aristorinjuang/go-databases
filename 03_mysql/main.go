package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	ID      int
	Title   string
	Content string
}

var (
	articles []article
	db       *sql.DB
	err      error
)

func index(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM articles;`)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		articles = append(articles, article{id, title, content})
	}

	response, err := json.Marshal(articles)
	if err != nil {
		log.Fatal(err)
	}

	if response != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var request article
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO articles(title, content) VALUES(?, ?)", request.Title, request.Content)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 0 {
		response, err := json.Marshal(article{int(rows), request.Title, request.Content})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func read(w http.ResponseWriter, r *http.Request, key int) {
	rows, err := db.Query(`SELECT * FROM articles WHERE id = ?`, key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id      int
		title   string
		content string
	)

	for rows.Next() {
		err = rows.Scan(&id, &title, &content)
		if err != nil {
			log.Fatal(err)
		}
	}

	response, err := json.Marshal(article{id, title, content})
	if err != nil {
		log.Fatal(err)
	}

	if title != "" && content != "" {
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func update(w http.ResponseWriter, r *http.Request, key int) {
	var request article
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE articles SET title = ?, content = ? WHERE id = ?", request.Title, request.Content, key)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 0 {
		response, err := json.Marshal(article{key, request.Title, request.Content})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func delete(w http.ResponseWriter, r *http.Request, key int) {
	result, err := db.Exec("DELETE FROM articles WHERE id = ?", key)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func main() {
	db, err = sql.Open("mysql", "root:root@/go_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := 0

		if len(r.URL.Path) > 1 {
			key, err = strconv.Atoi(strings.Split(r.URL.Path, "/")[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		switch r.Method {
		case http.MethodGet:
			if key == 0 {
				index(w, r)
			} else {
				read(w, r, key)
			}
		case http.MethodPost:
			create(w, r)
		case http.MethodPut:
			update(w, r, key)
		case http.MethodDelete:
			delete(w, r, key)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8000", nil)
}
