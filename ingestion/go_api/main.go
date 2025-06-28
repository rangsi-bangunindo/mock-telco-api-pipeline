package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/cdrs", cdrsHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Server running at http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
