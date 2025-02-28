package main

// Time to practice what you learned!

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line. - Done
	hobbies := [3]string{"Tennis", "Bicycle", "Software Development"}
	fmt.Printf("hobbies: %v, %T\n", hobbies, hobbies)
	// 2) Also output more data about that array:
	//   - The first element (standalone)
	//   - The second and third element combined as a new list
	fmt.Println("First element:", hobbies[:1])
	fmt.Println("Second and third:", hobbies[1:])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	var firstSecond []string
	firstSecond = append(firstSecond, hobbies[:2]...)
	fmt.Printf("Option 1 - First and second: %v, %T\n", firstSecond, firstSecond)

	firstSecond2 := hobbies[:2]
	fmt.Println("Option 2 - First and second:", firstSecond2)

	firstSecond3 := make([]string, 2)
	for i, v := range hobbies[:2] {
		firstSecond3[i] = v
	}
	fmt.Println("Option 3 - first and second:", firstSecond3)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	newSlice := hobbies[:]
	fmt.Printf("newSlice %v, len: %v, cap %v %T\n", newSlice, len(newSlice), cap(newSlice), newSlice)
	newSlice = newSlice[1:]
	fmt.Printf("newSlice second and third: %v, len: %v, cap %v %T\n", newSlice, len(newSlice), cap(newSlice), newSlice)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Master Go concepts", "Master Fullstack Develment in Go"}
	fmt.Println("Course goals:", courseGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Master Web Frontend with template and HTMX"
	courseGoals = append(courseGoals, "Master frontend with Vuejs and React")
	fmt.Println("Goals changes:", courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type truck struct {
		title string
		id    string
		price float64
	}

	sil1500 := truck{
		title: "Silverado 1500",
		id:    "S15002025",
		price: 45999.00,
	}

	f150 := truck{
		title: "Toyota Tundra",
		id:    "TT2025",
		price: 59995.00,
	}
	tundra := truck{
		title: "Ford F150",
		id:    "F1502025",
		price: 65999.00,
	}

	trucks := []truck{}
	trucks = append(trucks, sil1500, f150, tundra)

	fmt.Println("Truck:", trucks)

}
