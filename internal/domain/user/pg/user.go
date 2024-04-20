// Package pg contains implementations for user repositories.
package pg

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/ent/predicate"
	entuser "github.com/blackPavlin/shop/ent/user"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
)

// UserRepository pg repository implementation.
type UserRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewUserRepository create instance of UserRepository.
func NewUserRepository(client *ent.Client, logger *zap.Logger) *UserRepository {
	return &UserRepository{client: client, logger: logger}
}

// Create user in db.
func (r *UserRepository) Create(ctx context.Context, props *user.Props) (*user.User, error) {
	row, err := r.client.User.Create().
		SetName(props.Name).
		SetPhone(props.Phone).
		SetEmail(props.Email).
		SetRole(entuser.Role(user.RoleUser.String())).
		SetPassword(props.Password).
		Save(ctx)
	if err != nil {
		if pg.IsUniqueViolationErr(err) {
			return nil, errorx.ErrAlreadyExists
		}
		r.logger.Error("create user error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	created, err := mapDomainUserFromRow(row)
	if err != nil {
		r.logger.Error("convert user role error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return created, nil
}

// Get user from db.
func (r *UserRepository) Get(ctx context.Context, filter *user.Filter) (*user.User, error) {
	row, err := r.client.User.Query().
		Where(makeUserPredicate(filter)...).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get user error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	usr, err := mapDomainUserFromRow(row)
	if err != nil {
		r.logger.Error("convert user role error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return usr, nil
}

// Exist user in db.
func (r *UserRepository) Exist(ctx context.Context, filter *user.Filter) (bool, error) {
	exist, err := r.client.User.Query().
		Where(makeUserPredicate(filter)...).
		Exist(ctx)
	if err != nil {
		r.logger.Error("exist user error:", zap.Error(err))
		return false, errorx.ErrInternal
	}
	return exist, nil
}

func makeUserPredicate(filter *user.Filter) []predicate.User {
	predicates := make([]predicate.User, 0)
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entuser.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.ID.Neq) > 0 {
		predicates = append(predicates, entuser.IDNotIn(filter.ID.Neq.ToInt64()...))
	}
	if len(filter.Email.Eq) > 0 {
		predicates = append(predicates, entuser.EmailIn(filter.Email.Eq...))
	}
	if len(filter.Email.Neq) > 0 {
		predicates = append(predicates, entuser.EmailNotIn(filter.Email.Neq...))
	}
	return predicates
}

func mapDomainUserFromRow(row *ent.User) (*user.User, error) {
	role, err := user.RoleString(string(row.Role))
	if err != nil {
		return nil, fmt.Errorf("convertation user role error: %w", err)
	}
	return &user.User{
		ID:        user.ID(row.ID),
		Role:      role,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &user.Props{
			Email:    row.Email,
			Name:     row.Name,
			Phone:    row.Phone,
			Password: row.Password,
		},
	}, nil
}
