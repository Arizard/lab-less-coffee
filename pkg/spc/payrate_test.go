package spc

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewHourlyPayRateCoster(t *testing.T) {
	t.Run("Creates new payRateCosterHourly", func(t *testing.T) {
		coster := NewPayRateCosterHourly(decimal.NewFromFloat(18.50))

		asString := coster.String()
		expectedAsString := "18.50 per hour"
		if asString != expectedAsString {
			t.Errorf("coster.String() = %v, want %v", asString, expectedAsString)
		}
	})
}
