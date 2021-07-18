package spc

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type PayRateCurrency int

type PayRateCoster interface {
	GetCost(bl Block) decimal.Decimal
	String() (description string)
}

type payRateCosterHourly struct {
	rate decimal.Decimal
}

func (c *payRateCosterHourly) GetCost(bl Block) decimal.Decimal {

	dur := bl.End.Sub(bl.Start)

	hours := decimal.NewFromFloat(dur.Hours())

	return c.rate.Mul(hours)
}

func (c *payRateCosterHourly) String() (description string) {
	return fmt.Sprintf("%s per hour", c.rate.StringFixedBank(2))
}

func NewPayRateCosterHourly(rate decimal.Decimal) PayRateCoster {
	return &payRateCosterHourly{
		rate: rate,
	}
}

type PayRate struct {
	Name     string
	Metadata map[string]interface{}
	GetBlock PayRuleFunc
	Coster   PayRateCoster
}

func NewPayRate() *PayRate {
	payAll := All()
	return &PayRate{
		Name:     "Untitled Pay Rate",
		Coster:   NewPayRateCosterHourly(decimal.NewFromFloat(0)),
		GetBlock: payAll,
	}
}
