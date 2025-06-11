package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"` // "credit" or "debit"
	Timestamp int64   `json:"timestamp"`
}

// TransactionsClient provides methods to fetch transactions without calling HTTP manually
type TransactionsClient struct {
	Endpoint string
}

// NewTransactionsClient creates a new client instance
func NewTransactionsClient(endpoint string) *TransactionsClient {
	return &TransactionsClient{Endpoint: endpoint}
}

// GetTransactions fetches the list of transactions from the mock server
func (c *TransactionsClient) GetTransactions() ([]Transaction, error) {
	resp, err := http.Get(c.Endpoint + "/transactions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch transactions")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	if err := json.Unmarshal(body, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
