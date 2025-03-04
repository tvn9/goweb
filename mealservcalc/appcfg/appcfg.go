package appcfg

import (
	"errors"
	"mealservcalc/fileio"
	"mealservcalc/typeconv"
)

type FilePath struct {
	PriceFile  string // "./data/price.txt" is the price list
	TaxFile    string // "./data/taxrate.txt" is the tax rate list
	ResultFile string // "./data/results.txt"
}

var (
	taxList   []float64
	priceList []float64
)

func New() *FilePath {
	return &FilePath{}
}

func (f *FilePath) ChangePriceFilePath(s string) {
	f.PriceFile = s
}

func (f *FilePath) ChangeTaxFilePath(s string) {
	f.TaxFile = s
}

func (f *FilePath) RemoveAllFilePath() {
	f.PriceFile = ""
	f.TaxFile = ""
	f.ResultFile = ""
}

func (f FilePath) PriceFilePath() (string, error) {
	if f.PriceFile == "" {
		return f.PriceFile, errors.New("file path is empty")
	} else {
		return f.PriceFile, nil
	}
}

func (f FilePath) TaxFilePath() (string, error) {
	if f.TaxFile == "" {
		return f.TaxFile, errors.New("file path is empty")
	} else {
		return f.TaxFile, nil
	}
}

func (f FilePath) IsEmpty() bool {
	if f.PriceFile == "" || f.TaxFile == "" && f.ResultFile == "" {
		return true
	}
	return false
}

func (f FilePath) LoadTaxList() error {
	// taxRateList is a slice of string that stores tax rates
	taxRates := fileio.ReadDataFromFile(f.TaxFile)
	if len(taxRates) == 0 {
		return errors.New("error: tax rate list is empty")
	}
	for _, tr := range taxRates {
		taxList = append(taxList, typeconv.StringToFloat(tr))
	}
	return nil
}

func (f FilePath) LoadPriceList() error {
	// priceFile is a slice of string that store pricelist
	prices := fileio.ReadDataFromFile(f.PriceFile)
	if len(priceList) == 0 {
		return errors.New("error: price list is empty")
	}
	for _, ps := range prices {
		priceList = append(taxList, typeconv.StringToFloat(ps))
	}
	return nil
}
