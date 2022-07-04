package model

import (
	"strings"

	"github.com/Masterminds/squirrel"
)

type (
	Sort struct {
		column string
		order  string
	}
)

// NewSort Создание новой сортировки
func NewSort(column, order string) Sort {
	order = strings.ToUpper(order)
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}

	return Sort{
		column: column,
		order:  order,
	}
}

// UseSelectBuilder Добавление в squirrel.SelectBuilder сортировки
func (opt Sort) UseSelectBuilder(builder squirrel.SelectBuilder) squirrel.SelectBuilder {
	return builder.OrderBy(opt.column + " " + opt.order)
}
