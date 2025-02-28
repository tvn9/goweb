package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Println(result)

	result = recursionFactorial(5)
	fmt.Println(result)
}

func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
		fmt.Printf("%d: %d\n", i, result)
	}
	return result
}

func recursionFactorial(n int) int {
	i := 0
	if n == 0 {
		return 1
	}
	result := n * factorial(n-1)
	fmt.Printf("%d: %d\n", i+1, result)
	return result
}
