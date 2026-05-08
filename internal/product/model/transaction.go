package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CostDetails struct {
	Subtotal   int64 `json:"subtotal" bson:"subtotal"`
	Shipping   int64 `json:"shipping" bson:"shipping"`
	Tax        int64 `json:"tax" bson:"tax"`
	Discount   int64 `json:"discount" bson:"discount"`
	ServiceFee int64 `json:"service_fee" bson:"service_fee"`
}

type Transaction struct {
	ID         bson.ObjectID   `json:"id" bson:"_id,omitempty"`
	CreatedAt  time.Time       `json:"createdAt" bson:"createdAt"`
	CreatedBy  string          `json:"createdBy" bson:"createdBy"`
	Products   []bson.ObjectID `json:"products" bson:"products"`
	Costs      CostDetails     `json:"costs" bson:"costs"`
	GrandTotal int64           `json:"grand_total" bson:"grand_total"`
}
