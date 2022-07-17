package mongodb

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrFailidTypecastObjectID = errors.New("Failid typecast ObjectID to string")
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("users")

	return &UserRepository{
		collection: collection,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *entities.User) ([12]byte, error) {
	res, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return [12]byte{}, err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return [12]byte{}, ErrFailidTypecastObjectID
}

func (u *UserRepository) FindByID(ctx context.Context, id string) (*entities.User, error) {
	user := &entities.User{}

	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := u.collection.FindOne(ctx, bson.M{"_id": uid}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}

	if err := u.collection.FindOne(ctx, bson.M{"email": email}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
