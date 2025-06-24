package main

import "time"

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Region      string `json:"region"`
}

type CallLog struct {
	CallerID   int       `json:"caller_id"`
	ReceiverID int       `json:"receiver_id"`
	Timestamp  time.Time `json:"timestamp"`
	Duration   int       `json:"duration_sec"`
}

type DataUsage struct {
	UserID   int    `json:"user_id"`
	Date     string `json:"date"`
	DataUsed int    `json:"data_used_mb"`
}
