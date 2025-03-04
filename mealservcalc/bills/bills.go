package bills

import (
	"log"
	"mealservcalc/calc"
	"mealservcalc/items"
	"mealservcalc/typeconv"
)

type address struct {
	streetAddr string
	streetName string
	city       string
	state      string
	zipCode    string
}

type customer struct {
	fName string
	lName string
	phone string
	address
}

type table struct {
	id         int
	numsPeople int
	items      items.Product
}

type Bill struct {
	Item  map[string]float64
	Tax   float64
	Price float64
	Tip   float64
	Total float64
}

type Receipts struct {
	Id    string
	Title string
	customer
	table
	Bill
}

type RecRecords map[string][]*Receipts

func (b *Bill) ProcessBill(tip float64) {
	// bill is a struct that hold information about price, tax, tip, and total amount
	bill := New()

	tempBill := Bill{}
	if tempBill == bill {
		log.Fatal("bill data is not available")
	} else {
		for _, tr := range taxRateList {
			bill.Tax = typeconv.StringToFloat(tr)
		}

		for _, p := range priceList {
			bill.Price = typeconv.StringToFloat(p)
		}
	}

	billTotal, err := calc.MealCalc(bill.Price, bill.Tax, bill.Tip)
	if err != nil {
		log.Fatal(err)
	}

}

// func (j *PriceTax) calcPrices() {
// 	j.TaxRate = leio.ReadDataFromFile()
// 	result := make(map[string]float64)
// 	for _, p := range j.Price {
// 		priceStr := typeconv.FloatToString(p)
// 		result[priceStr] = p * (1 + j.TaxRate)
// 	}
// 	job.PriceNTax = result
// }

// New returns a new empty object
func New() *Bill {
	return &Bill{}
}

// Print
// func (job PriceTaxJob) Print() {
// 	fmt.Println("Tax rate:", typeconv.FloatToString(job.TaxRate))
// 	for k, v := range job.PriceNTax {
// 		fmt.Printf("Price: %s, Plus tax: %.2f\n", k, v)
// 	}
// 	fmt.Println()
// }

// func (job *PriceTaxJob) writePricesToFile(prices []string) {
// 	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		log.Fatalf("failed opening %s for writing: %s\n", name, err)
// 	}
// 	defer file.Close()

// 	for _, str := range prices {
// 		s := fmt.Sprintf("Tax rate: %s, price: %s, price & tax: %s", str, str, str)
// 		fmt.Println(s)
// 		if _, err := file.Write([]byte(s)); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
