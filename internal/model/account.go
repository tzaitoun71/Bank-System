package model

type Account struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
