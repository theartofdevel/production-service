package model

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"production_service/pkg/api/filter"
)

type Filterable interface {
	Enrich(query sq.SelectBuilder, alias string) sq.SelectBuilder
}

type filters struct {
	limit, offset uint64
	fields        []Field
}

func NewFilters(options filter.Filterable) *filters {
	var fs []Field
	var ff Field
	for _, f := range options.Fields() {
		ff = Field{
			Name:     f.Name,
			Operator: f.Operator,
			Value:    f.Value,
			Type:     f.Type,
		}
		fs = append(fs, ff)
	}

	return &filters{limit: options.Limit(), offset: options.Offset(), fields: fs}
}

func (f *filters) Enrich(query sq.SelectBuilder, alias string) sq.SelectBuilder {
	if len(f.fields) == 0 {
		return query
	}

	and := sq.And{}
	for _, where := range f.fields {
		var e sq.Sqlizer

		field := fmt.Sprintf("%s.%s", alias, where.Name)
		if alias == "" {
			field = where.Name
		}
		if where.Type == "date" {
			field = fmt.Sprintf("%s::date", field)
		}
		value := where.Value
		switch where.Operator {
		case "eq":
			e = sq.Eq{field: value}
		case "neq":
			e = sq.NotEq{field: value}
		case "like":
			e = sq.ILike{field: fmt.Sprintf("%%%s%%", value)}
		case "gt":
			e = sq.Gt{field: value}
		case "lt":
			e = sq.Lt{field: value}
		case "gte":
			e = sq.GtOrEq{field: value}
		case "lte":
			e = sq.LtOrEq{field: value}
		case "between":
			e = sq.Expr(fmt.Sprintf("'[%s]'::daterange @> %s", value, field))
		default:
			e = sq.Expr(fmt.Sprintf("%s %s %s", field, where.Operator, where.Value))
		}
		and = append(and, e)
	}

	query = query.Where(and)
	if f.limit == 0 {
		return query
	}
	return query.Limit(f.limit).Offset(f.offset)
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}
