package main

import "fmt"

func arrayTest() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a:", a)

	// get part of the array
	b := a[1:3]
	fmt.Println("index position 1 to 3 =", b)
	fmt.Printf("Len: %d, Cap: %d\n", len(a), cap(a))
	fmt.Printf("Len: %d, Cap: %d\n", len(b), cap(b))

	x := b
	x[0] = 99
	fmt.Println("x:", x)
	fmt.Printf("Len: %d, Cap: %d\n", len(x), cap(x))

	y := x
	y[1] = 999
	fmt.Println("a", a)
	fmt.Printf("z: %d\n", y)

	z := x[0:4]
	fmt.Println("z := x[0:4]", z)
	fmt.Println("a", a)

	j := a[:]
	fmt.Println("a", a)
	fmt.Println("j", j)

	c := a[1:]
	fmt.Println("index position 1 to 5 =", c)
	fmt.Printf("Len: %d, Cap: %d\n", len(a), cap(a))
	fmt.Printf("Len: %d, Cap: %d\n", len(c), cap(c))

	d := a[:]
	fmt.Println("index position 1 to 5 =", d)
	fmt.Printf("Len: %d, Cap: %d\n", len(d), cap(d))
	fmt.Printf("Len: %d, Cap: %d\n", len(d), cap(d))
}
