package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func serveJSON(w http.ResponseWriter, filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	serveJSON(w, "data/users.json")
}

func dataUsageHandler(w http.ResponseWriter, r *http.Request) {
	serveJSON(w, "data/data_usage.json")
}

func callLogsHandler(w http.ResponseWriter, r *http.Request) {
	serveJSON(w, "data/call_logs.json")
}

func main() {
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/data-usage", dataUsageHandler)
	http.HandleFunc("/api/call-logs", callLogsHandler)

	log.Println("Starting server on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
