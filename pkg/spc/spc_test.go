package spc

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestPayRate01(t *testing.T) {

	payRate := NewPayRate()
	payRate.Name = "Test Pay Rate 01"
	payRate.Coster = NewPayRateCosterHourly(decimal.NewFromFloat(17.55))

	payRate.GetBlock = DurationAfter(time.Hour*7 + time.Minute*36)

	shift := Block{
		Start: newTime(9, 0, 0),
		End:   newTime(18, 0, 0),
	}

	result := payRate.GetBlock(shift, shift.Start, shift.End)
	cost := payRate.Coster.GetCost(result)

	if !cost.Equal(decimal.NewFromFloat(24.57)) {
		t.Errorf("cost is %s, expected 24.57", cost.StringFixedBank(2))
	}
}
