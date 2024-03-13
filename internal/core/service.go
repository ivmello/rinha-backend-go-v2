package core

type Service interface {
	GetBalance(userID int) GetBalanceOutput
	CreateTransaction(input CreateTransactionInput) CreateTransactionOutput
}

type service struct {
	userRepository        UserRepository
	transactionRepository TransactionRepository
}

func NewService(userRepository UserRepository, transactionRepository TransactionRepository) Service {
	return &service{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
	}
}

func (s *service) GetBalance(userID int) GetBalanceOutput {
	return GetBalanceOutput{}
}

func (s *service) CreateTransaction(input CreateTransactionInput) CreateTransactionOutput {
	return CreateTransactionOutput{}
}
