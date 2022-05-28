package repository

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

// FindByEmail - Получение пользователя по email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	if err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

// Create - Создание нового пользователя
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	if _, err := r.collection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}
