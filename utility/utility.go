// Just some example of go package
package utility

import (
	"fmt"
	"time"
)

type USD int
type Type string

type Amount struct {
	Dollars USD
	Cents   USD
}

type USDBills struct {
	Cent1, Cent5, Cent10, Cent50, Usd1, Usd2, Usd5, Usd10,
	Usd20, Usd50, Usd100 USD
}

type Transaction struct {
	ID     int
	Amount Amount
	Payer  Party
	Payee  Party
	Date   string
}

type Address struct {
	StreetAddr, City, State, Zip, Country string
}

type SSNCard struct {
	SSN       string
	IssueDate time.Time
	Country   string
}
type DLCard struct {
	ID          string
	BD          string
	State       string
	Expired     string
	EyeColor    string
	Height      string
	Nationality string
	Sex         string
}

type Party struct {
	Name    string
	Contact Person
	Address Address
}

type Person struct {
	Fname, Lname string
	SNN          SSNCard
	DL           DLCard
	Phone        string
	Email        string
	Addr         Address
}

type Account struct {
	OpenDate    time.Time
	Type        Type
	Owner       Person
	AccNum      string
	OpenBalance Amount
	Income      Amount
	Expense     Amount
	EndBalance  Amount
	Activities  Transaction
}

// just return what ever string received
func Echo(s string) string {
	return s
}

// Add - adds two numbers of type int and return result of type int
func Add(n1, n2 int) int {
	return n1 + n2
}

// multiply - multiplys two numbers of type int and return result of type int
func Multiply(n1, n2 int) int {
	return n1 * n2
}

// Devide - devides two numbers of type int and return result of they int
func Devide(n1, n2 int) int {
	return n1 / n2
}

// Add money
func (b *USDBills) CountDollarsCents(d, c USD) (dollars, cents USD) {
	b.Cent1 = 1
	b.Usd1 = b.Cent1 * 100

	d = d * b.Usd1
	c = c * b.Cent1
	c = c + d

	dollars = c / b.Usd1
	cents = c % b.Usd1
	return dollars, cents
}

// AddBills allow input of bill in kinds starting from 1USD, 5USD, 10USD, 20USD, 50USD, and 100USD
func (b *USDBills) CountBills(oneUSD, twoUSD, fiveUSD, tenUSD, twentyUSD, fityUSD, hundredUSD USD) (usdTotal USD) {

	b.Usd1 = 1
	b.Usd2 = 2
	b.Usd5 = 5
	b.Usd10 = 10
	b.Usd20 = 20
	b.Usd50 = 50
	b.Usd100 = 100
	usd1Sub := oneUSD * b.Usd1
	usd2Sub := twoUSD * b.Usd2
	usd5Sub := fiveUSD * b.Usd5
	usd10Sub := tenUSD * b.Usd10
	usd20Sub := twentyUSD * b.Usd20
	usd50Sub := fityUSD * b.Usd50
	usd100Sub := hundredUSD * b.Usd100

	usdTotal = usd1Sub + usd2Sub + usd5Sub + usd10Sub + usd20Sub + usd50Sub + usd100Sub
	return usdTotal
}

// Adding Coins

func (a *Account) UpdateAcctBalance() *Account {
	a.EndBalance.Dollars = a.Income.Dollars - a.Expense.Dollars
	a.EndBalance.Cents = a.Income.Cents - a.Expense.Cents

	a.UpdateOpeningBalance()

	return a
}

// Deposit
func (a *Account) Deposit(d, c USD) {
	ds := d
	cs := c
	var dollars, cents USD
	var usd USDBills

	dollars, cents = usd.CountDollarsCents(ds, cs)

	// Update Transaction ID
	a.Activities.ID = a.UpdateTransacID()

	// Record Transaction Amount
	a.Activities.Amount.Dollars = dollars
	a.Activities.Amount.Cents = cents

	// Account Balance
	a.Income.Dollars += dollars
	a.Income.Cents += cents
	a.EndBalance.Dollars += a.Income.Dollars
	a.EndBalance.Cents += a.Income.Cents

	// Date Account Balance
	a.UpdateAcctBalance()

	// Record Transaction Time
	a.Activities.Date = a.GetDateTime()
}

func (a *Account) WithDraw(d, c USD) {
	ds := d
	cs := c
	var dollars, cents USD
	var usd USDBills

	// Update Transation ID
	a.Activities.ID = a.UpdateTransacID()

	dollars, cents = usd.CountDollarsCents(ds, cs)
	a.Expense.Dollars += dollars
	a.Expense.Cents += cents

	ds = a.Expense.Dollars
	cs = a.Expense.Cents

	a.Expense.Dollars, a.Expense.Cents = usd.CountDollarsCents(ds, cs)

	// Date Account Balance
	a.UpdateAcctBalance()

	// Mark Transaction date and time
	a.Activities.Date = a.GetDateTime()
}

// Update Transaction ID
func (a *Account) UpdateTransacID() int {
	if a.Activities.ID <= 0 {
		a.Activities.ID = 1
	} else {
		a.Activities.ID += 1
	}
	fmt.Println(a.Activities.ID)
	return a.Activities.ID
}

// Update Openning Balance
func (a *Account) UpdateOpeningBalance() {
	if a.Activities.ID == 1 {
		a.OpenBalance.Dollars = a.Activities.Amount.Dollars
		a.OpenBalance.Cents = a.Activities.Amount.Cents
	}
}

func (a *Account) GetDateTime() string {
	t := time.Now()
	// h := t.Local().Hour()
	// m := t.Local().Minute()
	// s := t.Local().Second()
	// y := t.Year()
	// mo := t.Month()
	// d := t.Date()

	// date = fmt.Sprintf("%v/%v/%v ", y, mo, d)
	// time = fmt.Sprintf("%v:%v:%d ", h, m, s)

	return t.String()

}
