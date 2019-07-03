package main

import "fmt"

func deferf() {
	fmt.Println("Defer!")
}

func main() {
	// run when this function exits
	defer deferf()
	fmt.Println("Hello, World!")

	// arrays:
	var numbers [5]int
	// initialized to zero
	fmt.Println("numbers:", numbers)

	// create and init, infer length:
	array := [...]int{1,2,3,4,5}
	fmt.Println("array:", array)
	// indexing
	fmt.Println("array[3]:", array[3])

	// slices
	slice := []int{1,2,3,4,5}
	slice2 := slice[1:3]
	fmt.Println("slice2:", slice2)
	fmt.Println("cap(slice2):", cap(slice2))
	slice3 := append(slice2, 10)	
	fmt.Println("slice3:", slice3)
	fmt.Println("slice:", slice)

	// maps

	m := map[string]int{}
	m["1"] = 1
	m["2"] = 2
	fmt.Println("m:", m)
}



