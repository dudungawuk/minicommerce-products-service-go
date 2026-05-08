package service

import (
	"product-service/internal/product/repository"
)

type TransactionService interface {
	// CreateTransaction(ctx context.Context, transaction dto.CreateTransactionRequest) error
}

type transactionService struct {
	t repository.TransactionRepository
	p repository.ProductRepository
}

func NewTransactionService(t repository.TransactionRepository, p repository.ProductRepository) TransactionService {
	return &transactionService{
		t: t,
		p: p,
	}
}

// func (s *transactionService) CreateTransaction(ctx context.Context, transactionDTO dto.CreateTransactionRequest) error {
// 	shipping := 7000
// 	tax := 0.1
// 	discount := 0.1

// 	c := model.CostDetails{
// 		Subtotal: ,
// 	}

// 	t := model.Transaction{
// 		CreatedAt: time.Now(),
// 		CreatedBy: transactionDTO.CreatedBy,
// 		Products:  transactionDTO.ListProducts,
// 	}

// 	return s.t.Insert(ctx, transaction)
// }
