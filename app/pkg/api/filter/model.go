package filter

import (
	"fmt"
	"strings"
)

type Operator string

const (
	DataTypeStr       = "string"
	DataTypeInt       = "int"
	DataTypeBool      = "bool"
	DataTypeDate      = "date"
	DataTypeArray     = "array"
	DataTypeTimeArray = "timeArray"
	DataTypeNull      = "empty"

	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "le"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "ge"
	OperatorIn            = "in"
	OperatorLike          = "like"
)

type Filterable interface {
	Limit() uint64
	Offset() uint64
	Fields() []Field
	AddFullField(rawValue string) error
	AddField(name string, operator Operator, value string) error
}

type Opts struct {
	filter      string
	limit       uint64
	offset      uint64
	filterTypes map[string]string
	fields      []Field
}

func NewOptions(limit, offset uint64, filterTypes map[string]string) *Opts {
	return &Opts{
		limit:       limit,
		offset:      offset,
		filterTypes: filterTypes,
	}
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

func (o *Opts) Limit() uint64 {
	return o.limit
}

func (o *Opts) Offset() uint64 {
	return o.offset
}

func (o *Opts) Fields() []Field {
	return o.fields
}

func (o *Opts) AddFullField(rawValue string) error {
	split := strings.Split(rawValue, " ")
	name := split[0]
	operator := split[1]
	value := split[2]

	err := validateOperator(operator)
	if err != nil {
		return err
	}
	dType, ok := o.filterTypes[name]
	if !ok {
		return fmt.Errorf("unknown param:`%s`", value)
	}
	if checkIsValueTuple(value) {
		dType = DataTypeArray
		if dType == DataTypeDate {
			dType = DataTypeTimeArray
		}
	}

	if (dType == DataTypeArray || dType == DataTypeTimeArray) && operator != OperatorIn {
		return fmt.Errorf("with array type name you can use only `in` operator. wrong query param:`%s=%s`",
			name, rawValue)
	}

	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: operator,
		Type:     dType,
	})
	return nil
}

func (o *Opts) AddField(name string, operator Operator, value string) error {
	err := validateOperator(string(operator))
	if err != nil {
		return err
	}
	dType, ok := o.filterTypes[name]
	if !ok {
		return fmt.Errorf("unknown param:`%s`", value)
	}
	if checkIsValueTuple(value) {
		dType = DataTypeArray
		if dType == DataTypeDate {
			dType = DataTypeTimeArray
		}
	}

	if (dType == DataTypeArray || dType == DataTypeTimeArray) && operator != OperatorIn {
		return fmt.Errorf("with array type name you can use only `in` operator. wrong query param:`%s, %s, %s`",
			name, operator, value)
	}

	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: string(operator),
		Type:     dType,
	})
	return nil
}

func checkIsValueTuple(value string) bool {
	split := strings.Split(value, ",")
	return len(split) != 1
}

func validateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorLike:
	case OperatorNotEq:
	case OperatorLowerThan:
	case OperatorLowerThanEq:
	case OperatorGreaterThan:
	case OperatorGreaterThanEq:
	case OperatorIn:
	default:
		return ErrBadOperator
	}
	return nil
}
