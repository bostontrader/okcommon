package utils

//{"min_fee":"0.00040000","currency":"BTC","max_fee":"0.01000000"}
type WithdrawalFee struct {
	CurrencyID string `json:"currency"`
	MinFee     string `json:"min_fee"`
	MaxFee     string `json:"max_fee"`
}
