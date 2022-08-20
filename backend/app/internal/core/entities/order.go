package entities

type OrderID [12]byte

type Order struct {
	ID       OrderID        `json:"id" bson:"_id,omitempty"`
	UserID   UserID         `json:"userId" bson:"user_id"`
	Address  Address        `json:"address" bson:"address"`
	Products []OrderProduct `json:"products" bson:"products"`
}

type OrderProduct struct {
	ProductID ProductID `json:"productId" bson:"product_id"`
	Amount    int64     `json:"amount" bson:"count"`
}
