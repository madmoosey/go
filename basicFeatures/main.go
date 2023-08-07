package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int())
	fmt.Println("Hello, Go")
	fmt.Println(20 + 20)
	fmt.Println(20 + 30)

	// Constants
	const price float32 = 275.00
	const tax float32 = 27.50
	const quantity = 2
	const base_price = price + tax
	const total_price = base_price * quantity
	fmt.Println(base_price)
	fmt.Println("Total:", total_price)

	// iota
	const (
		Watersports = iota
		Soccer
		Chess
	)
	fmt.Println(Watersports, Soccer, Chess)

	// variables
	var newprice float32 = 275.00
	var newtax float32 = 27.50
	fmt.Println(newprice + newtax)
	newprice = 300
	fmt.Println(newprice + newtax)

}

/*
Notes

int - hex, octal, or binary
uint - There are no uint literals. All literal whole numbers are treated as int values.
byte - There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run literals ('e') since the byte type is an alias for the uint8 type.
float64
bool
string
rune
*/
