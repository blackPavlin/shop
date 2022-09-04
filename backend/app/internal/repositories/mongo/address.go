package mongo

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/blackPavlin/shop/app/internal/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
func (a AddressRepository) Create(ctx context.Context, address *entities.Address) (entities.AddressID, error) {
	res, err := a.collection.InsertOne(ctx, address)
	if err != nil {
		return entities.AddressID{}, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return entities.AddressID(oid), nil
	}

	return entities.AddressID{}, errs.ErrFailidTypecastObjectID
}

// FindByID
func (a AddressRepository) FindByID(ctx context.Context, userID entities.UserID, id entities.AddressID) (*entities.Address, error) {
	address := &entities.Address{}

	if err := a.collection.FindOne(ctx, bson.M{"_id": id, "user_id": userID}).Decode(address); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrAddressNotFound
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

	if err := cursor.All(ctx, addresses); err != nil {
		return nil, err
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
