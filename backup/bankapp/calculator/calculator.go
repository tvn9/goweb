package calculator

import (
	"bankapp/database"
	"bankapp/reports"
	"fmt"
)

func Deposit(depositAmount float64) {
	balanceBeforeDeposit := database.ReadDataFromFile()
	b := balanceBeforeDeposit + depositAmount
	database.SaveDataToFile(b)
	reports.TransactionReport(depositAmount, 0.0, balanceBeforeDeposit)
}

func Withdraw(withdrawAmount float64) {
	balanceBeforeWithdraw := database.ReadDataFromFile()
	if balanceBeforeWithdraw <= 0.0 || balanceBeforeWithdraw < withdrawAmount {
		fmt.Println("Sorry, the account balance does not have enough money")
		fmt.Println("")
	} else {
		b := balanceBeforeWithdraw - withdrawAmount
		database.SaveDataToFile(b)
		reports.TransactionReport(0.0, withdrawAmount, balanceBeforeWithdraw)
	}
}
