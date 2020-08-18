package utils

type AccountsEntry struct {
	AccountID  string `json:"id"`
	Available  string `json:"available"`
	Balance    string `json:"balance"`
	CurrencyID string `json:"currency"`
	Frozen     string `json:"frozen"`
	Hold       string `json:"hold"`
	Holds      string `json:"holds"`
}

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
	To         int    `json:"to"`
}

type DepositHistory struct {
	Amount     string `json:"amount"`
	TXID       string `json:"txid"`
	CurrencyID string `json:"currency"`
	From       string `json:"from"`
	To         string `json:"to"`
	DepositID  int    `json:"deposit_id"`
	Timestamp  string `json:"timestamp"`
	Status     string `json:"status"`
}

type LedgerEntry struct {
    Amount    string
    Balance   string
    Currency  string
    Fee       string
    LedgerID  string `json:"ledger_id"`
    Typename  string
    Timestamp string `json:"timestamp"`
}

type WalletEntry struct {
	Available  string `json:"available"`
	Balance    string `json:"balance"`
	CurrencyID string `json:"currency"`
	Hold       string `json:"hold"`
}

type WithdrawalFee struct {
	CurrencyID string `json:"currency"`
	MinFee     string `json:"min_fee"`
	MaxFee     string `json:"max_fee"`
}
