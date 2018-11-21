package main

import (
	"github.com/shopspring/decimal"
)

// PercentageAdd adds to the price
func PercentageAdd(price decimal.Decimal, margin float64) decimal.Decimal {
	p, _ := price.Float64()
	res := p + (p * margin / 100)
	return decimal.NewFromFloat(res).Ceil()
}

// PercentageRemove removes from the price
func PercentageRemove(price decimal.Decimal, margin float64) decimal.Decimal {
	p, _ := price.Float64()
	res := p - (p * margin / 100)
	return decimal.NewFromFloat(res).Ceil()
}
