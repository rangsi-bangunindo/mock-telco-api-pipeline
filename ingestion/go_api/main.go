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
	http.HandleFunc("/api/call-logs", callLogsHandler)
	http.HandleFunc("/api/data-usage", dataUsageHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Starting server on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
