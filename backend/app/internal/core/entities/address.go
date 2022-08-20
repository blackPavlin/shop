package entities

type AddressID [12]byte

type Address struct {
	ID       AddressID `json:"id" bson:"_id,omitempty"`
	UserID   UserID    `json:"userId" bson:"user_id"`
	City     string    `json:"city" bson:"city"`
	Country  string    `json:"country" bson:"country"`
	Flat     *int      `json:"flat,omitempty" bson:"flat,omitempty"`
	House    int       `json:"house" bson:"house"`
	Letter   *string   `json:"letter,omitempty" bson:"letter,omitempty"`
	Postcode int       `json:"postcode" bson:"postcode"`
	Street   string    `json:"street" bson:"street"`
}
