package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ready\n"))
	})
	http.ListenAndServe(":18080", nil)
}
