package main

import (
	"fmt"
	"time"
)

type person struct {
	fname string
	lname string
	age   int
	bday  string
}

func sliceTest() {
	a := []string{"A", "B", "C"}
	fmt.Printf("a slice: %s, len: %d, cap: %d\n", a, len(a), cap(a))

	b := a
	fmt.Println("Assign (a) slice to (b) variable", a, b)

	z := b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("z:", z)

	z[0] = "ZZZZ"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("z:", z)

	b[0] = "AAAA"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("z:", z)

	z = append(z, b...)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("z:", z)

	b[0] = "ZZZZ"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("z:", z)

	typeAny := new(any)

	*typeAny = 1
	fmt.Printf("typeAny: %T %v\n", *typeAny, *typeAny)

	*typeAny = "Thanh Nguyen"
	fmt.Printf("typeAny: %T %v\n", *typeAny, *typeAny)

	*typeAny = 129.99
	fmt.Printf("typeAny: %T %v\n", *typeAny, *typeAny)

	*typeAny = person{
		fname: "Thanh",
		lname: "Nguyen",
		age:   55,
		bday:  time.Now().Format(time.DateOnly),
	}
	fmt.Printf("typeAnyA: %T %v\n", *typeAny, *typeAny)
}
