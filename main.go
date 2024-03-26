package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Title string
	Body string
	id int
}

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		
		p := Post{"Post", "Body", 1}

		jsonData, err := json.Marshal(p)

		if err != nil {
			w.Write([]byte("Error"))
		}

		w.Write([]byte(jsonData))
	})

	http.ListenAndServe(":3000", nil)
}