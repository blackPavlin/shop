package entities

type ProductID [12]byte

type Product struct {
	ID          ProductID  `json:"id" bson:"_id,omitempty"`
	Name        string     `json:"name" bson:"name"`
	Description string     `json:"description" bson:"description"`
	CategoryID  CategoryID `json:"categoryId" bson:"category_id"`
	Amount      int64      `json:"amount" bson:"amount"`
}
