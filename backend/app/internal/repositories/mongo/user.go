package mongo

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/core/entities"
	"github.com/blackPavlin/shop/app/internal/core/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrFailidTypecastObjectID = errors.New("Failid typecast ObjectID")
)

// UserRepository
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository
func NewUserRepository(db *mongo.Database) UserRepository {
	collection := db.Collection("users")

	return UserRepository{
		collection: collection,
	}
}

// Create
func (u UserRepository) Create(ctx context.Context, user *entities.User) (entities.UserID, error) {
	res, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return entities.UserID{}, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return entities.UserID(oid), nil
	}

	return entities.UserID{}, ErrFailidTypecastObjectID
}

// FindByID
func (u UserRepository) FindByID(ctx context.Context, id entities.UserID) (*entities.User, error) {
	user := &entities.User{}

	if err := u.collection.FindOne(ctx, bson.M{"_id": id}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

// FindByEmail
func (u UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}

	if err := u.collection.FindOne(ctx, bson.M{"email": email}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errs.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

// ExistsByEmail
func (u UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	count, err := u.collection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}