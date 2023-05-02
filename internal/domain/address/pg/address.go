package pg

import (
	"context"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entaddress "github.com/blackPavlin/shop/ent/address"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
)

// AddressRepository
type AddressRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewAddressRepository
func NewAddressRepository(client *ent.Client, logger *zap.Logger) *AddressRepository {
	return &AddressRepository{client: client, logger: logger}
}

// Create
func (r *AddressRepository) Create(
	ctx context.Context,
	props *address.Props,
) (*address.Address, error) {
	user, ok := user.GetUser(ctx)
	if !ok {
		return nil, errorx.ErrUnauthorized
	}

	row, err := r.client.Address.Create().
		SetUserID(int64(user.ID)).
		SetCountry(props.Country).
		SetCity(props.City).
		SetFlat(*props.Flat). // TODO: WTF???
		SetHouse(props.House).
		SetLetter(*props.Letter). // TODO: WTF???
		SetPostcode(props.Postcode).
		SetStreet(props.Street).
		Save(ctx)
	if err != nil {
		if pg.IsForeignKeyViolationErr(err, "address_user_fk") {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("create address error", zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return mapDomainAddressFromRow(row), nil
}

// Get
func (r *AddressRepository) Get(
	ctx context.Context,
	filter *address.Filter,
) (*address.Address, error) {
	predicate := makePredicates(filter)
	// TODO: Get UserID from ctx

	row, err := r.client.Address.Query().
		Where(predicate...).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get address error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return mapDomainAddressFromRow(row), nil
}

// Query
func (r *AddressRepository) Query(
	ctx context.Context,
	criteria *address.QueryCriteria,
) (address.Addresses, error) {
	predicate := makePredicates(criteria.Filter)
	// TODO: Get UserID from ctx

	rows, err := r.client.Address.Query().
		Where(predicate...).
		All(ctx)
	if err != nil {
		r.logger.Error("query address error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return mapDomainAddressesFromRows(rows), nil
}

func makePredicates(filter *address.Filter) []predicate.Address {
	predicates := make([]predicate.Address, 0)
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entaddress.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.UserID.Eq) > 0 {
		predicates = append(predicates, entaddress.UserIDIn(filter.UserID.Eq.ToInt64()...))
	}
	return predicates
}

func mapDomainAddressFromRow(row *ent.Address) *address.Address {
	return &address.Address{
		ID:     address.ID(row.ID),
		UserID: user.ID(row.UserID),
		Props: &address.Props{
			City:     row.City,
			Country:  row.Country,
			Flat:     &row.Flat, // TODO: WTF???
			House:    row.House,
			Letter:   &row.Letter, // TODO: WTF???
			Postcode: row.Postcode,
			Street:   row.Street,
		},
	}
}

func mapDomainAddressesFromRows(rows ent.Addresses) address.Addresses {
	result := make(address.Addresses, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainAddressFromRow(row))
	}
	return result
}
