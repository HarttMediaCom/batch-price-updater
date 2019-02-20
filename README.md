# Shopify Batch Price Updater

[![Build Status](https://travis-ci.org/HarttMediaCom/batch-price-updater.svg?branch=master)](https://travis-ci.org/HarttMediaCom/batch-price-updater)
[![Go Report Card](https://goreportcard.com/badge/github.com/HarttMediaCom/batch-price-updater)](https://goreportcard.com/report/github.com/HarttMediaCom/batch-price-updater)
[![GoDoc](https://godoc.org/github.com/HarttMediaCom/batch-price-updater?status.svg)](https://godoc.org/github.com/HarttMediaCom/batch-price-updater)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

Batch update cost of items or product price by margin

## How it works
Depending on margin set, this tool can batch update the prices ( product or cost of item).

If argument `-cost` is added, it will take the product price and lower the price by margin and set the field of cost of item.

Example: If product cost $150 and margin is set to 20, the cost of item will be set at $120 (150-20%)

If argument `-product` is added, it will take the cost of item and add margin to it.

Example: If cost of item is $150 and margin is set to 20, the product price will be set at $180 (150+20%)

If there is no cost of item set, the tool will skip the calculation.

## Requirements
* Go compiler
* Generated Shopify credentials - [private app](https://help.shopify.com/en/manual/apps/private-apps)
* Permissions: read_products, write_products, read_inventory and write_inventory

## Building
```sh
go build
```

## Configuration
### Copy `.env.example` to `.env`

| Value  | Type  | Default  | Description  |
|---|---|---|---|
| APP_KEY  | string  | none  | Shopify API Key  |
| APP_PASSWORD  | string  | none  | Shopify API password  |
| SHOP_NAME  | string  | none  |  Shop name without .myshopify.com |
| MARGIN  | float  | 20  | Margin in percetange eg 25, 14.50...  |

## Running
To update cost of items by margin
```sh
./batch-price-updater -cost
```

To update product price by cost of item + margin
```sh
./batch-price-updater -product
```

## License
MIT

## Author
Hartt Media 
