package main

import (
	"net/http"
	"strconv"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, users)
}

func cdrsHandler(w http.ResponseWriter, r *http.Request) {
	// Optional query param: ?limit=20
	limit := 20
	if l := r.URL.Query().Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	writeJSON(w, generateCDRs(limit))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
