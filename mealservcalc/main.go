package main

import (
	"fmt"
	"log"
	"mealservcalc/appcfg"
)

func init() {
	filePath := appcfg.New()
	filePath.PriceFile = "./data/price.txt"
	filePath.TaxFile = "./data/taxrate.txt"
	filePath.ResultFile = "./data/results.txt"

	// Load price list from file to []string
	if err := filePath.LoadPriceList(); err != nil {
		log.Fatalf("failed to load price list to memory - %v", err)
	}
	if err := filePath.LoadTaxList(); err != nil {
		log.Fatalf("failed to load tax list to memory - %v", err)
	}

}

func main() {
	app := *appcfg.New()
	fmt.Println(app)

}
