package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.Split(r.URL.Path, "/")[1]

		if key != "" {
			switch r.Method {
			case http.MethodGet:
				file, err := ioutil.ReadFile("data/" + key + ".json")
				if err != nil {
					log.Println(err)
				}
				if file != nil {
					w.Header().Set("Content-Type", "application/json")
					w.Write(file)
				}
			case http.MethodPost:
				body, err := io.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}
				if body != nil {
					err = ioutil.WriteFile("data/"+key+".json", body, 0644)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	})
	http.ListenAndServe(":80", nil)
}
