package reports

import (
	"bankapp/database"
	"fmt"
)

func TransactionReport(d, w, bb float64) {
	fmt.Println()
	fmt.Printf("Account blance: %.2f\n", bb)
	if d > 0 {
		fmt.Printf("Deposit amount: %.2f\n", d)
		fmt.Printf("Balance after depost is: %.2f\n", database.ReadDataFromFile())
	}
	if w > 0 {
		fmt.Printf("Withdraw amount: %.2f\n", w)
		fmt.Printf("Balance after withdraw: %.2f\n", database.ReadDataFromFile())
	}
}
