package repository

import (
	"context"
	"fmt"
	"product-service/internal/product/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TransactionRepository interface {
	Insert(ctx context.Context, transaction model.Transaction) error
}

type transactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) TransactionRepository {
	return &transactionRepository{
		collection: db.Collection("transaction"),
	}
}

func (r *transactionRepository) Insert(ctx context.Context, transaction model.Transaction) error {
	result, err := r.collection.InsertOne(ctx, transaction)

	if err != nil {
		return err
	}

	fmt.Println("success add with id ", result.InsertedID)

	return nil
}
