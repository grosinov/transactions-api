package dtos

type Balance struct {
	Balance     float64 `json:"balance"`
	TotalDebit  uint    `json:"total_debit"`
	TotalCredit uint    `json:"total_credit"`
}
