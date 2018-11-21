package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestPercentage(t *testing.T) {

	add := PercentageAdd(decimal.NewFromFloat(150.00), 20.00)
	if !add.Equal(decimal.NewFromFloat(180)) {
		t.Error("calculation wrong, expect 180 got ", add.String())
		return
	}

	remove := PercentageRemove(decimal.NewFromFloat(150.00), 20.00)
	if !remove.Equal(decimal.NewFromFloat(120)) {
		t.Error("calculation wrong, expect 120 got ", add.String())
	}

}

func TestService(t *testing.T) {
	service := NewService(
		"",
		"",
		"",
	)

	err := service.UpdateCostOfItems(0.00)
	if err == nil {
		t.Error("UpdateCostOfItems we should get an error, got nil")
		return
	}

	if err.Error() != "margin cannot be 0 or less" {
		t.Error("We should get an error of margin but got something else")
		return
	}

	err = service.UpdateProductPrices(2.00)
	if err == nil {
		t.Error("UpdateProductPrices we should get an error, got nil")
		return
	}

	// this should failed before calling callback
	err = service.run(func(a int, b int, d *decimal.Decimal, margin float64) error {
		return nil
	}, 20.00)

	if err == nil {
		t.Error("run we should get an error, got nil")
		return
	}
}
