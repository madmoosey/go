package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	/*IMPORTANT: Omittint a type does not have the same effect for variables as it does
	for constants.
		var price = 275.00 - The compiler will always infer the type of literal floating-point values as float64,
		var tax float32 = 27.50
		fmt.Println(price + tax)
		invalid operation: price + tax (mismatched types float64 and float32)
	*/

	var newprice float32 = 275.00
	var newtax float32 = 27.50
	fmt.Println(newprice + newtax)
	newprice = 300
	fmt.Println(newprice + newtax)
	var newprice2 = 375.00
	var newprice3 = newprice2
	fmt.Println(newprice3)

	// Variables can be defined without an initial value
	var blankprice float32
	fmt.Println(blankprice)
	blankprice = 475.00
	fmt.Println(blankprice)

	// multiple vars
	var mprice, mtax = 275.00, 27.50
	fmt.Println(mprice + mtax)

	// walrus declaration
	// The short variable declaration syntax can be used only within functions
	// Go doesn’t usually allow variables to be redefined but makes a limited exception when the short syntax is used
	wprice := 465.00
	fmt.Println(wprice)
	wprice2, wtax, winStock := 275.00, 27.50, true
	fmt.Println("Total:", wprice2+wtax)
	fmt.Println("In stock:", winStock)

	wprice3, wtax := 547.50, 54.75
	fmt.Println(wprice3 + wtax)

	// Blank identifier
	// It is illegal in Go to define a variable and not use it
	var _ = "Alice"

	// Pointers
	first := 100
	second := 100

	first++

	fmt.Println(first)
	fmt.Println(second)

	var third *int = &first // Pointers are defined by the ampersand (address operator)

	first++
	*third++

	fmt.Println("First:", first)
	fmt.Println("Third:", third)  // Memory location
	fmt.Println("Third:", *third) // Dereferenced

	var fourth *int = third
	*fourth++

	fmt.Println("First:", first)
	fmt.Println("Third:", *third)

	var fifth *int
	fmt.Println(fifth)
	// fmt.Println(*fifth) This throws ans error because the value is not assigned
	fifth = &first
	fmt.Println(*fifth)

	sixth := &fifth
	fmt.Println(**sixth)

	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]
	fmt.Println(secondName)
	sort.Strings(names[:])
	fmt.Println(secondName)

	newNames := [3]string{"Alice", "Charlie", "Bob"}
	secondPosition := &newNames[1]
	fmt.Println(*secondPosition)
	sort.Strings(newNames[:])
	fmt.Println(*secondPosition)
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
