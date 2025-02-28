package main

import (
	"fmt"
	"math"
)

func main() {
	var inflationRate float64
	var investAmt float64

	var years float64
	expectReturn := 5.5

	fmt.Print("Enter amount: ")
	fmt.Scan(&investAmt)

	fmt.Print("Enter number of year: ")
	fmt.Scan(&years)

	fmt.Print("Expected return: ")
	fmt.Scan(&expectReturn)

	fmt.Print("Enter inflation rate: ")
	fmt.Scan(&inflationRate)

	// futureValue := investAmt * math.Pow(1+expectReturn/100, years)
	fv := futureValueCalc(investAmt, expectReturn, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	rfv := futureRealValueCalc(fv, inflationRate, years)

	fmt.Println("future value:", fv)
	fmt.Println("future real value:", rfv)
}

func futureValueCalc(i, e, y float64) (f float64) {
	return i * math.Pow(1+e/100, y)
}

func futureRealValueCalc(f, i, y float64) float64 {
	return f / math.Pow(1+i/100, y)
}
