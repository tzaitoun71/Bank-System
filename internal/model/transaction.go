package model

import "time"

type Transaction struct {
	TxID      int64     `json:"tx_id"`
	AccountID int64     `json:"account_id"`
	Amount    float64   `json:"amount"`
	TxType    string    `json:"tx_type"`
	Timestamp time.Time `json:"timestamp"`
}
