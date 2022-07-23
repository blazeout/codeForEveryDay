package main

import (
	"log"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", pong)
	log.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
