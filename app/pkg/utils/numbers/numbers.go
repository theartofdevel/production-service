package numbers

import (
	"github.com/shopspring/decimal"
)

var (
	One     = decimal.NewFromInt(1)
	Hundred = decimal.NewFromInt(100)
)

// PercentOfNumber calculates amount as percent of amount
// deposit amount = 200
// processingFeePercent = 2
// 200 - (200*(1-2/100) = 200 - (196) = 4 = -4
func PercentOfNumber(percent decimal.Decimal, amount decimal.Decimal) decimal.Decimal {
	return amount.Sub(amount.Mul(One.Sub(percent.Div(Hundred))))
}
