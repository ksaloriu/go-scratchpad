package main

import (
	"fmt"
	"github.com/ksaloriu/go-scratchpad/util"
)

func deferf() {
	fmt.Println("Defer!")
}

func return_values() (string, int, string) {
	return "1", 2, "3"
}

type user struct {
	name  string
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

	// vars
	var x int
	x = 1
	var y = 2
	z := 3
	fmt.Println("x:", x, " y:", y, " z:", z)

	// functions
	n1, n2, n3 := return_values()
	fmt.Println("n1:", n1, " n2:", n2, " n3:", n3)
	_, n4, _ := return_values()
	fmt.Println("n4:", n4)

	// imports
	fmt.Println("Exported_var:", util.Exported_var)

	// arrays
	var numbers [5]int
	// initialized to zero
	fmt.Println("numbers:", numbers)
	// create and init, infer length:
	array := [...]int{1, 2, 3, 4, 5}
	fmt.Println("array:", array)
	// indexing
	fmt.Println("array[3]:", array[3])

	// slices
	slice := []int{1, 2, 3, 4, 5}
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

	// pointers
	value := 13
	var p1 *int
	p2 := &value
	p1 = &value
	fmt.Println("*p1:", *p1, " *p2:", *p2)
	*p1 = 14
	fmt.Println("*p1:", *p1, " *p2:", *p2)

	// methods
	user2.print()
	user2.changeEmail("foo@bar.com")
	user2.print()
}
