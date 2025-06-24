package main

import "net/http"

func usersHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, users)
}

func callLogsHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, generateCallLogs())
}

func dataUsageHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, generateDataUsage())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
