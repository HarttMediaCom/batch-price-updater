package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/bold-commerce/go-shopify"
)

type Client struct {
	shopify *goshopify.Client
}

type InventoryItemRequest struct {
	Item InventoryItem `json:"inventory_item"`
}

type InventoryItem struct {
	ID   int             `json:"id"`
	Cost decimal.Decimal `json:"cost"`
}

const (
	defaultSleepTime = 10
)

// AllProducts fetches all available products from shopify store
func (c *Client) AllProducts() ([]goshopify.Product, error) {
	options := struct {
		Page  int `url:"page"`
		Limit int `url:"limit"`
	}{1, 250}

	var allProducts []goshopify.Product
	apiCounter := 1
	for {
		products, err := c.shopify.Product.List(options)
		if err != nil {
			return []goshopify.Product{}, err
		}

		if len(products) <= 0 {
			break
		}

		allProducts = append(allProducts, products...)
		options.Page++

		if apiCounter > 35 {
			time.Sleep(time.Duration(defaultSleepTime) * time.Second)
			apiCounter = 1
		}

	}

	return allProducts, nil
}

// UpdateCost calls inventory API for update
func (c *Client) UpdateCost(id int, price decimal.Decimal) error {
	var req InventoryItemRequest
	req.Item.ID = id
	req.Item.Cost = price
	return c.shopify.Put(fmt.Sprintf("/admin/inventory_items/%d.json", id), req, nil)
}

// ProductCost calls inventory API to get cost
func (c *Client) ProductCost(invID int) (decimal.Decimal, error) {
	var req InventoryItemRequest
	err := c.shopify.Get(fmt.Sprintf("/admin/inventory_items/%d.json", invID), &req, nil)
	return req.Item.Cost, err
}
