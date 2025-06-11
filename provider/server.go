package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dragonator/coach-ai-assignment/provider/client"
)

func main() {
	http.HandleFunc("/transactions", transactionsHandler)
	log.Println("Mock transaction server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func transactionsHandler(w http.ResponseWriter, r *http.Request) {
	transactions := generateMockTransactions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// generateMockTransactions returns a list of mock transactions with some inconsistencies
func generateMockTransactions() []client.Transaction {
	staticTransactions := []client.Transaction{
		{ID: "txn-12345", UserID: "user-1", Amount: 100.50, Type: "credit", Timestamp: time.Now().Add(-10 * time.Minute).Unix()},
		{ID: "txn-67890", UserID: "user-2", Amount: 75.25, Type: "debit", Timestamp: time.Now().Add(-5 * time.Minute).Unix()},
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5) + 5 // Generate between 5-10 transactions
	transactions := make([]client.Transaction, n)

	for i := 0; i < n; i++ {
		transactions[i] = client.Transaction{
			ID:        randomID(),
			UserID:    randomID(),
			Amount:    float64(rand.Intn(10000)) / 100.0,
			Type:      randomType(),
			Timestamp: time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Second).Unix(),
		}
	}

	// Introduce occasional duplicates
	if rand.Intn(10) > 7 {
		transactions = append(transactions, transactions[rand.Intn(len(transactions))])
	}

	// Introduce missing fields
	if len(transactions) > 0 && rand.Intn(10) > 7 {
		transactions[0].UserID = ""
	}

	return append(transactions, staticTransactions...)
}

func randomID() string {
	return time.Now().Format("20060102150405") + string(rand.Intn(100))
}

func randomType() string {
	types := []string{"credit", "debit"}
	return types[rand.Intn(len(types))]
}
