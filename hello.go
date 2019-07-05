package main

import "fmt"

func deferf() {
	fmt.Println("Defer!")
}

type user struct {
	name string
	email string
}

// method with value receiver
func (u user) print() {
	fmt.Printf("User: %s<%s>\n", u.name, u.email)
}

// method with pointer receiver
func (u *user) changeEmail(email string) {
	u.email = email
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

	// structs
	var user1 user
	user2 := user{"Nobody", "nobody@example.com"}
	fmt.Println("user1:", user1)
	fmt.Println("user2:", user2)

	// methods
	user2.print()
	user2.changeEmail("foo@bar.com")
	user2.print()
}



