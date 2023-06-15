package numbers

import (
	"testing"

	"github.com/shopspring/decimal"
)

/*
	deposit amount = 200
	processingFeePercent = 2
	200 - (200*(1-2/100) = 200 - (196) = 4 = -4
*/

type testCase struct {
	Amount, Percent, Result decimal.Decimal
}

func newTestCase(amount, percent, result float64) *testCase {
	return &testCase{
		Amount:  decimal.NewFromFloat(amount),
		Percent: decimal.NewFromFloat(percent),
		Result:  decimal.NewFromFloat(result),
	}
}

func TestPercentOfNumber(t *testing.T) {
	cases := []*testCase{
		newTestCase(200, 0, 0),
		newTestCase(200, 2, 4),
		newTestCase(900, 35, 315),
		newTestCase(900, 3.5, 31.5),
		newTestCase(800.499, 67.4, 539.536326),
		newTestCase(8030400, 15.4, 1236681.6),
		newTestCase(0.00005, 0.1030, 0.0000000515),
	}
	for _, c := range cases {
		res := PercentOfNumber(c.Percent, c.Amount)
		if !res.Equal(c.Result) {
			t.Fatalf("expect %s, got %s", c.Result.String(), res.String())
		}
	}
}
