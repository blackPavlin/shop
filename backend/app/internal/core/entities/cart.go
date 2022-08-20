package entities

type CartID [12]byte

type Cart struct {
	ID       CartID        `json:"id" bson:"_id,omitempty"`
	UserID   UserID        `json:"userId" bson:"user_id"`
	Products []CartProduct `json:"products" bson:"products"`
}

type CartProduct struct {
	ProductID ProductID `json:"productId" bson:"product_id"`
	Amount    int64     `json:"amount" bson:"amount"`
}
