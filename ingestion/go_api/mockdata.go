package main

import (
	"math/rand"
	"time"
)

var users = []User{
	{ID: 1, Name: "Alice", PhoneNumber: "081234567890", Region: "Jakarta"},
	{ID: 2, Name: "Bob", PhoneNumber: "081298765432", Region: "Bandung"},
	{ID: 3, Name: "Charlie", PhoneNumber: "081356789012", Region: "Surabaya"},
}

func generateCallLogs() []CallLog {
	var logs []CallLog
	for i := 0; i < 10; i++ {
		caller := users[rand.Intn(len(users))].ID
		receiver := users[rand.Intn(len(users))].ID
		for caller == receiver {
			receiver = users[rand.Intn(len(users))].ID
		}
		logs = append(logs, CallLog{
			CallerID:   caller,
			ReceiverID: receiver,
			Timestamp:  time.Now().Add(-time.Duration(rand.Intn(72)) * time.Hour),
			Duration:   rand.Intn(500) + 30,
		})
	}
	return logs
}

func generateDataUsage() []DataUsage {
	var usage []DataUsage
	for _, user := range users {
		usage = append(usage, DataUsage{
			UserID:   user.ID,
			Date:     time.Now().Format("2006-01-02"),
			DataUsed: rand.Intn(5000) + 100,
		})
	}
	return usage
}
