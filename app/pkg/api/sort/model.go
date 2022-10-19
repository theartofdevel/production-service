package sort

import (
	"strings"
)

type Order string

const (
	OrderASC  Order = "ASC"
	OrderDESC Order = "DESC"
)

type Sortable interface {
	Field() string
	Order() string
}

type opts struct {
	field string
	order string
}

func NewOptions(field string) *opts {
	sortOrder := OrderASC
	if strings.HasPrefix(field, "-") {
		sortOrder = OrderDESC
		field = strings.TrimPrefix(field, "-")
	}
	return &opts{
		field: field,
		order: string(sortOrder),
	}
}

func (o *opts) Field() string {
	return o.field
}
func (o *opts) Order() string {
	return o.order
}
