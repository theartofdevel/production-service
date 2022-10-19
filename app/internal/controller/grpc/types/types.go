package types

import (
	pb_common_filter "github.com/theartofdevel/production-service-contracts/gen/go/common/filter/v1"
	"production_service/pkg/api/filter"
)

type FilterOperator = filter.Operator

func IntOperatorFromPB(e pb_common_filter.IntFieldFilter_Operator) FilterOperator {
	switch e {
	case pb_common_filter.IntFieldFilter_OPERATOR_EQ:
		return filter.OperatorEq
	case pb_common_filter.IntFieldFilter_OPERATOR_NEQ:
		return filter.OperatorNotEq
	case pb_common_filter.IntFieldFilter_OPERATOR_LT:
		return filter.OperatorLowerThan
	case pb_common_filter.IntFieldFilter_OPERATOR_LTE:
		return filter.OperatorLowerThanEq
	case pb_common_filter.IntFieldFilter_OPERATOR_GT:
		return filter.OperatorGreaterThan
	case pb_common_filter.IntFieldFilter_OPERATOR_GTE:
		return filter.OperatorGreaterThanEq
	default:
		return ""
	}
}

func StringOperatorFromPB(e pb_common_filter.StringFieldFilter_Operator) FilterOperator {
	switch e {
	case pb_common_filter.StringFieldFilter_OPERATOR_EQ:
		return filter.OperatorEq
	case pb_common_filter.StringFieldFilter_OPERATOR_NEQ:
		return filter.OperatorNotEq
	case pb_common_filter.StringFieldFilter_OPERATOR_LIKE:
		return filter.OperatorLike
	default:
		return ""
	}
}
