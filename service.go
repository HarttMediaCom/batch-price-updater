package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/bold-commerce/go-shopify"
	log "github.com/sirupsen/logrus"
)

// Service contains runnable shopify client
type Service struct {
	shopifyClient *Client
	apiCounter    int
}

// NewService creates new service for shopify
func NewService(key, password, name string) *Service {
	app := goshopify.App{
		ApiKey:   key,
		Password: password,
	}

	return &Service{
		shopifyClient: &Client{goshopify.NewClient(app, name, "")},
	}
}

// UpdateCostOfItems updates cost of item by removing margin from product price
func (s *Service) UpdateCostOfItems(margin float64) error {
	log.Info("Updating cost of items...")
	if margin <= 0.0 {
		return fmt.Errorf("margin cannot be 0 or less")
	}

	err := s.run(s.costOfItem, margin)
	if err != nil {
		return err
	}
	log.Info("Service done")
	return nil
}

// UpdateProductPrices updates product price by adding margin on top of the cost of item
func (s *Service) UpdateProductPrices(margin float64) error {
	log.Info("Updating product prices...")
	if margin <= 0.0 {
		return fmt.Errorf("margin cannot be 0 or less")
	}
	err := s.run(s.productPrice, margin)
	if err != nil {
		return err
	}
	log.Info("Service done")
	return nil
}

func (s *Service) run(callback func(int, int, *decimal.Decimal, float64) error, margin float64) error {
	// fetch all products
	products, err := s.shopifyClient.AllProducts()
	if err != nil {
		return fmt.Errorf("error fetching all products, err: %s", err.Error())
	}

	for _, product := range products {
		for _, variant := range product.Variants {

			log.Info("Updating ", product.Title, " variant: ", variant.Title)

			err := callback(variant.ID, variant.InventoryItemId, variant.Price, margin)
			if err != nil {
				log.Error("error updating ", product.Title, " err: ", err)
			}

			if s.apiCounter > 35 {
				s.apiCounter = 1
				time.Sleep(defaultSleepTime * time.Second)
			}
		}
	}
	return nil
}

func (s *Service) costOfItem(variantID int, invID int, price *decimal.Decimal, margin float64) error {
	if !price.GreaterThan(decimal.NewFromFloat(0.0)) {
		return fmt.Errorf("price is not set")
	}
	s.apiCounter++
	newPrice := PercentageRemove(*price, margin)
	return s.shopifyClient.UpdateCost(invID, newPrice)
}

func (s *Service) productPrice(variantID int, invID int, price *decimal.Decimal, margin float64) error {
	s.apiCounter += 2
	cost, err := s.shopifyClient.ProductCost(invID)
	if err != nil {
		return err
	}
	if !cost.GreaterThan(decimal.NewFromFloat(0.0)) {
		return fmt.Errorf("cost of item not set")
	}

	newProductPrice := PercentageAdd(cost, margin)

	_, err = s.shopifyClient.shopify.Variant.Update(goshopify.Variant{
		ID:    variantID,
		Price: &newProductPrice,
	})

	return err
}
