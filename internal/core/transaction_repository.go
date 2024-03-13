package core

import (
	"database/sql"
)

type TransactionRepository interface {
	GetAll(userID int) []Transaction
	Create(payload Transaction)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) GetAll(userID int) []Transaction {
	return nil
}

func (t *transactionRepository) Create(payload Transaction) {
}
