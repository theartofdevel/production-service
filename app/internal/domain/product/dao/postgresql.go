package dao

import (
	"context"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"production_service/internal/dal"
	"production_service/internal/dal/postgres"
	"production_service/internal/domain/product/model"
	psql "production_service/pkg/postgresql"
	"production_service/pkg/tracing"
)

type ProductDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewProductStorage(client psql.Client) *ProductDAO {
	return &ProductDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *ProductDAO) All(ctx context.Context) ([]model.Product, error) {
	all, err := repo.findBy(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.Product, len(all))
	for i, e := range all {
		resp[i] = e.ToDomain()
	}

	return resp, nil
}

// Create
func (repo *ProductDAO) Create(ctx context.Context, req model.CreateProduct) error {
	sql, args, err := repo.qb.
		Insert(postgres.ProductTable).
		Columns(
			"id",
			"name",
			"description",
			"image_id",
			"price",
			"currency_id",
			"rating",
			"category_id",
			"specification",
			"created_at",
		).
		Values(
			req.ID,
			req.Name,
			req.Description,
			req.ImageID,
			req.Price,
			req.CurrencyID,
			req.Rating,
			req.CategoryID,
			req.Specification,
			req.CreatedAt,
		).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}

	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := repo.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return dal.ErrNothingInserted
	}

	return nil
}

func (repo *ProductDAO) findBy(ctx context.Context) ([]ProductStorage, error) {
	statement := repo.qb.
		Select(
			"id",
			"name",
			"description",
			"image_id",
			"price",
			"currency_id",
			"rating",
			"category_id",
			"specification",
			"created_at",
			"updated_at",
		).
		From(postgres.ProductTable + " p")

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	entities := make([]ProductStorage, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e ProductStorage
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Description,
			&e.ImageID,
			&e.Price,
			&e.CurrencyID,
			&e.Rating,
			&e.CategoryID,
			&e.Specification,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			tracing.Error(ctx, err)

			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}
