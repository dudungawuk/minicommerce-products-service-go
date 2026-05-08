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
	GetByID(ctx context.Context, id bson.ObjectID) (model.Product, error)
	// UserTransaction(ctx context.Context) ()
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database, id bson.ObjectID) ProductRepository {
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

	fmt.Println("DEBUG REPOSITORY: ", err)

	return products, nil

}

func (r *productRepository) GetByID(ctx context.Context, id bson.ObjectID) (model.Product, error) {
	var product model.Product

	// 1. Gunakan filter { "_id": id }
	// 2. Gunakan .Decode() untuk memasukkan data ke variabel product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (r *productRepository) Insert(ctx context.Context, product model.Product) error {
	result, err := r.collection.InsertOne(ctx, product)

	if err != nil {
		return err
	}

	fmt.Println("success add with id ", result.InsertedID)

	return nil

}
