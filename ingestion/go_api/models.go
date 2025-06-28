package main

import "time"

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Region      string `json:"region"`
}

type CDR struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Caller      string    `json:"caller"`
	Receiver    string    `json:"receiver"`
	Duration    int       `json:"duration_seconds"`
	CallType    string    `json:"call_type"`
	CellTowerID string    `json:"cell_tower_id"`
	Status      string    `json:"status"`
}
