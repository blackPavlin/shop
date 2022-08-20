package mongo

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrAddressNotFound        = errors.New("Address not found")
)

// AddressRepository
type AddressRepository struct {
	collection *mongo.Collection
}

// NewAddressRepository
func NewAddressRepository(db *mongo.Database) AddressRepository {
	collection := db.Collection("addresses")

	return AddressRepository{
		collection: collection,
	}
}

// Create
func (a AddressRepository) Create(ctx context.Context, address *entities.Address) (*entities.Address, error) {
	res, err := a.collection.InsertOne(ctx, address)
	if err != nil {
		return nil, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		address.ID = entities.AddressID(oid)
		return address, nil
	}

	return nil, ErrFailidTypecastObjectID
}

// FindByID
func (a AddressRepository) FindByID(ctx context.Context, userID entities.UserID, id entities.AddressID) (*entities.Address, error) {
	address := &entities.Address{}

	if err := a.collection.FindOne(ctx, bson.M{"_id": id, "user_id": userID}).Decode(address); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrAddressNotFound
		}
		return nil, err
	}

	return address, nil
}

// FindByUserID
func (a AddressRepository) FindByUserID(ctx context.Context, userID entities.UserID) ([]*entities.Address, error) {
	cursor, err := a.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	addresses := []*entities.Address{}

	for cursor.Next(ctx) {
		address := &entities.Address{}
		if err := cursor.Decode(address); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

// DeleteByID
func (a AddressRepository) DeleteByID(ctx context.Context, userID entities.UserID, id entities.AddressID) error {
	if _, err := a.collection.DeleteOne(ctx, bson.M{"_id": id, "user_id": userID}); err != nil {
		return err
	}

	return nil
}
