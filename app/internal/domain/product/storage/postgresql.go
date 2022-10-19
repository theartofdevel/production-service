package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"production_service/pkg/api/filter"
	"production_service/pkg/api/sort"
	db "production_service/pkg/client/postgresql/model"
	"production_service/pkg/logging"
)

type ProductStorage struct {
	queryBuilder sq.StatementBuilderType
	client       PostgreSQLClient
}

func NewProductStorage(client PostgreSQLClient) *ProductStorage {
	return &ProductStorage{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:       client,
	}
}

const (
	scheme = "public"
	table  = "product"
)

func (s *ProductStorage) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]Product, error) {
	sortDB := db.NewSortOptions(sorting)
	filterDB := db.NewFilters(filtering)

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

	query = filterDB.Enrich(query, "")
	query = sortDB.Sort(query, "")

	sql, args, err := query.ToSql()
	logger := logging.WithFields(ctx, map[string]interface{}{
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

	list := make([]Product, 0)

	for rows.Next() {
		p := Product{}
		if err = rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.ImageID, &p.Price, &p.CurrencyID, &p.Rating, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			err = db.ErrScan(err)
			logger.Error(err)
			return nil, err
		}

		list = append(list, p)
	}

	return list, nil
}
