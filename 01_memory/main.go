package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	data := make(map[string][]byte)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.Split(r.URL.Path, "/")[1]

		if key != "" {
			switch r.Method {
			case http.MethodGet:
				w.Header().Set("Content-Type", "application/json")
				w.Write(data[key])
			case http.MethodPost:
				body, err := io.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}
				if body != nil {
					data[key] = body
				}
			}
		}
	})
	http.ListenAndServe(":80", nil)
}
