package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var users = []User{
	{ID: 1, Name: "Alice", PhoneNumber: "+628123456789", Region: "Jakarta"},
	{ID: 2, Name: "Bob", PhoneNumber: "+628579876543", Region: "Bandung"},
	{ID: 3, Name: "Charlie", PhoneNumber: "+628998877665", Region: "Surabaya"},
	{ID: 4, Name: "Diana", PhoneNumber: "+628112233445", Region: "Medan"},
	{ID: 5, Name: "Evan", PhoneNumber: "+628667788990", Region: "Semarang"},
}

var callTypes = []string{"incoming", "outgoing"}
var callStatuses = []string{"completed", "missed", "failed"}
var cellTowers = []string{"CT-1029", "CT-1031", "CT-1100", "CT-1112", "CT-1205"}

func generateCDRs(n int) []CDR {
	var records []CDR
	for i := 0; i < n; i++ {
		caller := users[rand.Intn(len(users))].PhoneNumber
		receiver := users[rand.Intn(len(users))].PhoneNumber
		for caller == receiver {
			receiver = users[rand.Intn(len(users))].PhoneNumber
		}

		record := CDR{
			ID:          uuid.New().String(),
			Timestamp:   time.Now().Add(-time.Duration(rand.Intn(72)) * time.Hour),
			Caller:      caller,
			Receiver:    receiver,
			Duration:    rand.Intn(600) + 10, // seconds
			CallType:    callTypes[rand.Intn(len(callTypes))],
			CellTowerID: cellTowers[rand.Intn(len(cellTowers))],
			Status:      callStatuses[rand.Intn(len(callStatuses))],
		}
		records = append(records, record)
	}
	return records
}
