package pg

import (
	"context"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/ent/predicate"
	userent "github.com/blackPavlin/shop/ent/user"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
	"go.uber.org/zap"
)

// UserRepository
type UserRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewUserRepository
func NewUserRepository(client *ent.Client, logger *zap.Logger) *UserRepository {
	return &UserRepository{client: client, logger: logger}
}

// Create
func (r *UserRepository) Create(ctx context.Context, props *user.Props) (*user.User, error) {
	row, err := r.client.User.Create().
		SetName(props.Name).
		SetPhone(props.Phone).
		SetEmail(props.Email).
		SetRole(userent.Role(user.RoleUser.String())).
		SetPassword(props.Password).
		Save(ctx)
	if err != nil {
		if pg.IsUniqueViolationErr(err) {
			return nil, errorx.ErrAllreadyExists
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

// Get
func (r *UserRepository) Get(ctx context.Context, filter *user.Filter) (*user.User, error) {
	predicates := makeUserPredicate(&user.QueryCriteria{Filter: filter})

	row, err := r.client.User.Query().
		Where(predicates...).
		First(ctx)
	if err != nil {
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

// Exist
func (r *UserRepository) Exist(ctx context.Context, filter *user.Filter) (bool, error) {
	predicates := makeUserPredicate(&user.QueryCriteria{Filter: filter})

	exist, err := r.client.User.Query().
		Where(predicates...).
		Exist(ctx)
	if err != nil {
		r.logger.Error("exist user error:", zap.Error(err))
		return false, errorx.ErrInternal
	}

	return exist, nil
}

func makeUserPredicate(criteria *user.QueryCriteria) []predicate.User {
	predicates := make([]predicate.User, 0)

	return predicates
}

func mapDomainUserFromRow(row *ent.User) (*user.User, error) {
	role, err := user.UserRoleString(string(row.Role))
	if err != nil {
		return nil, err
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