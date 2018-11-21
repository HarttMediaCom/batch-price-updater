package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var costOfItem = flag.Bool("cost", false, "Updates cost of items")
var productPrice = flag.Bool("product", false, "Updates product price")

func main() {
	flag.Parse()

	if !*costOfItem && !*productPrice {
		log.Error("You need to provide a param which field you want to update. cost or product")
		flag.Usage()
		return
	}
	if err := godotenv.Load(); err != nil {
		log.Error(err)
		return
	}
	service := NewService(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_PASSWORD"),
		os.Getenv("SHOP_NAME"),
	)

	margin, err := strconv.ParseFloat(os.Getenv("MARGIN"), 64)
	if err != nil {
		log.Error("error parsing margin value, err: ", err)
		return
	}

	switch {
	case *costOfItem:
		service.UpdateCostOfItems(margin)
		break
	case *productPrice:
		service.UpdateProductPrices(margin)
		break
	}

}
