package mongo

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/blackPavlin/shop/app/internal/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// AddressRepository
type AddressRepository struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

// NewAddressRepository
func NewAddressRepository(db *mongo.Database, logger *zap.Logger) AddressRepository {
	collection := db.Collection("addresses")

	return AddressRepository{
		collection: collection,
		logger:     logger,
	}
}

// Create
func (a AddressRepository) Create(ctx context.Context, address *entities.Address) (entities.AddressID, error) {
	res, err := a.collection.InsertOne(ctx, address)
	if err != nil {
		a.logger.Error("failid create address", zap.Error(err))
		return entities.AddressID{}, errs.ErrInternal
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		a.logger.Error("failid typecast address id", zap.Error(errs.ErrFailidTypecastObjectID))
		return entities.AddressID{}, errs.ErrInternal
	}
	return entities.AddressID(oid), nil
}

// FindByID
func (a AddressRepository) FindByID(ctx context.Context, userID entities.UserID, id entities.AddressID) (*entities.Address, error) {
	address := &entities.Address{}
	if err := a.collection.FindOne(ctx, bson.M{
		"_id":     primitive.ObjectID(id),
		"user_id": primitive.ObjectID(userID),
	}).Decode(address); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrAddressNotFound
		}
		a.logger.Error("failid find address by id", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return address, nil
}

// FindByUserID
func (a AddressRepository) FindByUserID(ctx context.Context, userID entities.UserID) (entities.Addresses, error) {
	cursor, err := a.collection.Find(ctx, bson.M{"user_id": primitive.ObjectID(userID)})
	if err != nil {
		a.logger.Error("failid find address by user id", zap.Error(err))
		return nil, errs.ErrInternal
	}
	defer cursor.Close(ctx)
	addresses := &entities.Addresses{}
	if err := cursor.All(ctx, addresses); err != nil {
		a.logger.Error("failid find address by user id", zap.Error(err))
		return nil, errs.ErrInternal
	}
	return *addresses, nil
}

// DeleteByID
func (a AddressRepository) DeleteByID(ctx context.Context, userID entities.UserID, id entities.AddressID) error {
	if _, err := a.collection.DeleteOne(ctx, bson.M{
		"_id":     primitive.ObjectID(id),
		"user_id": primitive.ObjectID(userID),
	}); err != nil {
		a.logger.Error("failid delete address by id", zap.Error(err))
		return errs.ErrInternal
	}
	return nil
}
