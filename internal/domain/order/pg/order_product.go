package pg

import (
	"context"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entorderproduct "github.com/blackPavlin/shop/ent/orderproduct"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/order"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
)

// OrderProductRepository pg repository implementation.
type OrderProductRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewOrderProductRepository create instance of OrderProductRepository.
func NewOrderProductRepository(client *ent.Client, logger *zap.Logger) *OrderProductRepository {
	return &OrderProductRepository{client: client, logger: logger}
}

// BulkCreateTx order products in db with transaction.
func (r *OrderProductRepository) BulkCreateTx(ctx context.Context, products order.Products) (order.Products, error) {
	if len(products) == 0 {
		return nil, nil
	}
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		r.logger.Error("using tx in non tx context", zap.Error(errorx.ErrInternal))
		return nil, errorx.ErrInternal
	}
	client := tx.Client()
	rows, err := client.OrderProduct.
		CreateBulk(createOrderProductBuilders(client, products)...).
		Save(ctx)
	if err != nil {
		if pg.IsForeignKeyViolationErr(err, "order_product_order_fk") {
			return nil, errorx.ErrNotFound
		}
		if pg.IsForeignKeyViolationErr(err, "order_product_product_fk") {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("bulkCreate order products error", zap.Error(err))
		return nil, errorx.ErrUnauthorized
	}
	return mapDomainOrderProductsFromRows(rows), nil
}

func makeOrderProductPredicates(filter order.ProductFilter) []predicate.OrderProduct {
	predicates := make([]predicate.OrderProduct, 0)
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entorderproduct.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.ID.Neq) > 0 {
		predicates = append(predicates, entorderproduct.IDNotIn(filter.ID.Neq.ToInt64()...))
	}
	if len(filter.OrderID.Eq) > 0 {
		predicates = append(predicates, entorderproduct.OrderIDIn(filter.OrderID.Eq.ToInt64()...))
	}
	if len(filter.OrderID.Neq) > 0 {
		predicates = append(predicates, entorderproduct.OrderIDNotIn(filter.OrderID.Eq.ToInt64()...))
	}
	if len(filter.ProductID.Eq) > 0 {
		predicates = append(predicates, entorderproduct.ProductIDIn(filter.ProductID.Eq.ToInt64()...))
	}
	if len(filter.ProductID.Neq) > 0 {
		predicates = append(predicates, entorderproduct.ProductIDNotIn(filter.ProductID.Neq.ToInt64()...))
	}
	return predicates
}

func createOrderProductBuilder(client *ent.Client, pt *order.Product) *ent.OrderProductCreate {
	return client.OrderProduct.Create().
		SetOrderID(pt.OrderID.ToInt64()).
		SetProductID(pt.ProductID.ToInt64()).
		SetAmount(pt.Props.Amount).
		SetPrice(pt.Props.Price)
}

func createOrderProductBuilders(client *ent.Client, products order.Products) []*ent.OrderProductCreate {
	builders := make([]*ent.OrderProductCreate, 0, len(products))
	for _, prod := range products {
		builders = append(builders, createOrderProductBuilder(client, prod))
	}
	return builders
}

func mapDomainOrderProductFromRow(row *ent.OrderProduct) *order.Product {
	return &order.Product{
		ID:        order.ProductID(row.ID),
		OrderID:   order.ID(row.OrderID),
		ProductID: product.ID(row.ProductID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &order.ProductProps{
			Amount: row.Amount,
			Price:  row.Price,
		},
	}
}

func mapDomainOrderProductsFromRows(rows ent.OrderProducts) order.Products {
	result := make(order.Products, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainOrderProductFromRow(row))
	}
	return result
}
