package repository

import (
	"context"
	"fmt"
	"product-service/internal/product/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository interface {
	Insert(ctx context.Context, product model.Product) error
	FindAll(ctx context.Context) ([]model.Product, error)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepository{
		collection: db.Collection("products"),
	}
}

func (r *productRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil

}

func (r *productRepository) Insert(ctx context.Context, product model.Product) error {
	result, err := r.collection.InsertOne(ctx, product)

	if err != nil {
		return err
	}

	fmt.Println("success add with id ", result.InsertedID)

	return nil

}
