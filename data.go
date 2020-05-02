package utils

// {"name":"0x","currency":"ZRX","can_withdraw":"1","can_deposit":"1","min_withdrawal":"1.0000000000000000"}
type CurrenciesEntry struct {
	CanDeposit    string `json:"can_deposit"`
	CanWithdraw   string `json:"can_withdraw"`
	CurrencyID    string `json:"currency"`
	Name          string `json:"name"`
	MinWithdrawal string `json:"min_withdrawal"`
}

type DepositAddress struct {
	Address    string `json:"address"`
	Tag        string `json:"tag"`
	PaymentID  string `json:"payment_id"`
	Memo       string `json:"memo"`
	CurrencyID string `json:"currency"`
	To         string `json:"to"`
}

//{"balance":"1.00000000","available":"1.00000000","currency":"BSV","hold":"0.00000000"}
type WalletEntry struct {
	Available  string `json:"available"`
	Balance    string `json:"balance"`
	CurrencyID string `json:"currency"`
	Hold       string `json:"hold"`
}

//{"min_fee":"0.00040000","currency":"BTC","max_fee":"0.01000000"}
type WithdrawalFee struct {
	CurrencyID string `json:"currency"`
	MinFee     string `json:"min_fee"`
	MaxFee     string `json:"max_fee"`
}
