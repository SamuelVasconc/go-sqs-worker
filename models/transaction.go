package models

type Transaction struct {
	ID          int64   `json:"id"`
	Date        string  `json:"date"`
	Amount      float64 `json:"amount"`
	Observation string  `json:"obs"`
	Protocol    string  `json:"protocol"`
	Status      string  `json:"status"`
}
