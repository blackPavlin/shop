package entities

type CategoryID [12]byte

type Category struct {
	ID   CategoryID `json:"id" bson:"_id,omitempty"`
	Name string     `json:"name" bson:"name"`
}
