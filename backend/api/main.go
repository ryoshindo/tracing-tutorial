package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ready"))
	})
	http.ListenAndServe(":8080", nil)
}
