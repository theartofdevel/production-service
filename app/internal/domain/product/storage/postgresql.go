package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"production_service/internal/domain/product/model"
	"production_service/pkg/client/postgresql"
	db "production_service/pkg/client/postgresql/model"
	"production_service/pkg/logging"
)

type ProductStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewProductStorage(client PostgreSQLClient) ProductStorage {
	return ProductStorage{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

const (
	scheme = "public"
	table  = "product"
)

func (s *ProductStorage) All(ctx context.Context) ([]model.Product, error) {
	query := s.queryBuilder.Select("id").
		Column("name").
		Column("description").
		Column("image_id").
		Column("price").
		Column("currency_id").
		Column("rating").
		Column("created_at").
		Column("updated_at").
		From(scheme + "." + table)

	// TODO Задача №2*. Реализовать фильтрацию и сортировку по полям
	/*

		!!!! НЕ ДЕЛАТЬ
		Transport Layer: HTTP / AMQP / WS
		/api/products?name=eq:купон&price=lt:300&sort_by=created_at&sort_order=desc

		|
		V

		!!!! НЕ ДЕЛАТЬ
		Service Layer
		* FilterOptions --> FilterOptions for Storage
		* SortOptions --> SortOptions for Storage

		|
		V

		!!!! ДЕЛАТЬ !!!!
		Storage Layer

		1. Создать структуры сортировки и фильтрации которые будут аргументами в методе All
		2. Методы которые принимают query и обогащают его филтрацией и соритровкой.
		3. вызвать эти методы в методе product.storage.postgresql.All()

	*/

	sql, args, err := query.ToSql()
	logger := logging.GetLogger(ctx).WithFields(map[string]interface{}{
		"sql":   sql,
		"table": table,
		"args":  args,
	})
	if err != nil {
		err = db.ErrCreateQuery(err)
		logger.Error(err)
		return nil, err
	}

	logger.Trace("do query")
	rows, err := s.client.Query(ctx, sql, args...)
	if err != nil {
		err = db.ErrDoQuery(err)
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	list := make([]model.Product, 0)

	for rows.Next() {
		p := model.Product{}
		if err = rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.ImageID, &p.Price, &p.CurrencyID, &p.Rating, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			err = db.ErrScan(postgresql.ParsePgError(err))
			logger.Error(err)
			return nil, err
		}

		list = append(list, p)
	}

	return list, nil
}
