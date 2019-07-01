package main

import "fmt"

func deferf() {
	fmt.Println("Defer!")
}

func main() {
	// run when this function exits
	defer deferf()
	fmt.Println("Hello, World!")
}
