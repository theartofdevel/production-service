package model

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"production_service/pkg/api/sort"
)

type Sortable interface {
	Sort(query sq.SelectBuilder, alias string) sq.SelectBuilder
}

type sorts struct {
	field string
	order string
}

func NewSortOptions(options sort.Sortable) *sorts {
	return &sorts{
		field: options.Field(),
		order: options.Order(),
	}
}

func (s *sorts) Sort(query sq.SelectBuilder, alias string) sq.SelectBuilder {
	if s.field == "" {
		return query
	}
	field := s.field
	if alias != "" {
		field = fmt.Sprintf("%s.%s", alias, field)
	}
	return query.OrderBy(fmt.Sprintf("%s %s", field, s.order))
}
