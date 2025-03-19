package models

type FundTransferResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type BalanceResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"error,omitempty"`
	Balance float64 `json:"balance"`
}
