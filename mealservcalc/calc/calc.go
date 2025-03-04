package calc

import "errors"

func MealCalc(price, tax, tip float64) (total float64, err error) {
	if !amountInputCheck(price, tax, tip) {
		return total, errors.New("invalit amount entered")
	} else {
		taxSub := taxTotal(price, tax)
		tipSub := tipTotal(price, tip)
		total = price + taxSub + tipSub
		return total, nil
	}
}

func amountInputCheck(price, tax, tip float64) bool {
	if price < 0 || tax < 0 || tip < 0 {
		return false
	}
	return true
}

func taxTotal(price, tax float64) (total float64) {
	tax = tax / 100
	total = price * (1 + tax)
	return total
}

func tipTotal(price, tip float64) (total float64) {
	tip = tip / 100
	total = price * (1 + tip)
	return total
}
