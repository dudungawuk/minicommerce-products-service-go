package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type CreateTransactionRequest struct {
	CreatedBy    string          `json:"createdBy" binding:"required"`
	ListProducts []bson.ObjectID `json:"ListProducts" binding:"required"`
}

type CostDetailsResponse struct {
	Subtotal   int64 `json:"subtotal" `
	Shipping   int64 `json:"shipping" `
	Tax        int64 `json:"tax" `
	Discount   int64 `json:"discount" `
	ServiceFee int64 `json:"service_fee"`
}

type CreateTransactionResponse struct {
	ID         string              `json:"id"`
	CreatedAt  string              `json:"createdAt" `
	CreatedBy  string              `json:"createdBy" `
	Products   []bson.ObjectID     `json:"products" `
	Costs      CostDetailsResponse `json:"costs" `
	GrandTotal int64               `json:"grand_total"`
}
