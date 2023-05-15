package pg

import (
	"context"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entcategory "github.com/blackPavlin/shop/ent/category"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
)

// CategoryRepository
type CategoryRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewCategoryRepository
func NewCategoryRepository(client *ent.Client, logger *zap.Logger) *CategoryRepository {
	return &CategoryRepository{client: client, logger: logger}
}

// Create
func (r *CategoryRepository) Create(
	ctx context.Context,
	props *category.Props,
) (*category.Category, error) {
	row, err := r.client.Category.Create().
		SetName(props.Name).
		Save(ctx)
	if err != nil {
		if pg.IsUniqueViolationErr(err) {
			return nil, errorx.ErrAlreadyExists
		}

		r.logger.Error("create category error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainCategoryFromRow(row), nil
}

// Query
func (r *CategoryRepository) Query(
	ctx context.Context,
	criteria *category.QueryCriteria,
) (category.Categories, error) {
	predicates := makePredicate(criteria)
	rows, err := r.client.Category.
		Query().
		Where(predicates...).
		All(ctx)
	if err != nil {
		r.logger.Error("query categories error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainCategoriesFromRows(rows), nil
}

// Get
func (r *CategoryRepository) Get(
	ctx context.Context,
	filter *category.Filter,
) (*category.Category, error) {
	predicates := makePredicate(&category.QueryCriteria{Filter: *filter})
	row, err := r.client.Category.Query().
		Where(predicates...).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get category error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainCategoryFromRow(row), err
}

// Update
func (r *CategoryRepository) Update(
	ctx context.Context,
	id category.ID,
	props *category.Props,
) (*category.Category, error) {
	row, err := r.client.Category.
		UpdateOneID(int64(id)).
		SetName(props.Name).
		Save(ctx)
	if err != nil {
		r.logger.Error("update category error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainCategoryFromRow(row), err
}

// Delete
func (r *CategoryRepository) Delete(ctx context.Context, id category.ID) error {
	err := r.client.Category.
		DeleteOneID(int64(id)).
		Exec(ctx)
	if err != nil {
		r.logger.Error("delete category error", zap.Error(err))
		return errorx.ErrInternal
	}
	return nil
}

func makePredicate(criteria *category.QueryCriteria) []predicate.Category {
	predicates := make([]predicate.Category, 0)
	if len(criteria.Filter.ID.Eq) != 0 {
		predicates = append(predicates, entcategory.IDIn(criteria.Filter.ID.Eq.ToInt64()...))
	}
	if len(criteria.Filter.ID.Neq) != 0 {
		predicates = append(predicates, entcategory.IDNotIn(criteria.Filter.ID.Neq.ToInt64()...))
	}
	if len(criteria.Filter.Name.Eq) != 0 {
		predicates = append(predicates, entcategory.NameIn(criteria.Filter.Name.Eq...))
	}
	if len(criteria.Filter.Name.Neq) != 0 {
		predicates = append(predicates, entcategory.NameNotIn(criteria.Filter.Name.Neq...))
	}
	return predicates
}

func mapDomainCategoryFromRow(row *ent.Category) *category.Category {
	return &category.Category{
		ID:        category.ID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &category.Props{
			Name: row.Name,
		},
	}
}

func mapDomainCategoriesFromRows(rows ent.Categories) category.Categories {
	result := make(category.Categories, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainCategoryFromRow(row))
	}
	return result
}
