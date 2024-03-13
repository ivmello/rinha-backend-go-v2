package core

import "time"

type GetBalanceOutput struct {
	Balance struct {
		Total       int       `json:"total"`
		BalanceDate time.Time `json:"data_extrato"`
		Limit       int       `json:"limite"`
	} `json:"saldo"`
	Transactions []struct {
		Value       int       `json:"valor"`
		Operation   string    `json:"tipo"`
		Description string    `json:"descricao"`
		CreatedAt   time.Time `json:"realizada_em"`
	} `json:"ultimas_transacoes"`
}

type CreateTransactionInput struct {
	Value       int    `json:"valor"`
	Operation   string `json:"tipo"`
	Description string `json:"descricao"`
}

type CreateTransactionOutput struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}
