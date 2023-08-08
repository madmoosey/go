package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, Operations")

	// Arithmetic operators
	price, tax := 275.00, 27.40
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)

	// Go allows integer values to overflow by wrapping around, rather than reporting an error. Floating-point values overflow to positive or negative infinity
	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(intVal)
	fmt.Println(intVal * 2)
	fmt.Println(floatVal)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))

	// Remainder operator
	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))

	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)

	// Increment operator

	value := 102
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)

	// Concatenating Strings
	fmt.Println("########################## Concatenating Strings ##########################")
	greeting := "Hello"
	language := "Go"
	combinedString := greeting + ", " + language

	fmt.Println(combinedString)

	// Comparison Operators
	fmt.Println("########################## Comparison Operators ##########################")
	first := 100
	const second = 200.00 // The untyped constant is a floating-point value but can be represented as an integer value because its fractional digits are zeros

	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lessThan)
	fmt.Println(lessThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)

	// Go doesn’t provide a ternary operator, which means that expressions like this cannot be used

	var max int
	if first > second {
		max = first
	} else {
		max = second
	}

	fmt.Println(max)

	// Comparing Pointers
	fmt.Println("########################## Comparing Pointers ##########################")
	// Pointers can be compared to see if they point at the same memory location

	third := &first
	fourth := &first

	alpha := 100
	beta := &alpha

	fmt.Println(third == fourth)
	fmt.Println(third == beta)

	fmt.Println("Dereferenced Variables")
	fmt.Println(*third == *fourth)
	fmt.Println(*third == *beta)

	// Logical operators
	fmt.Println("########################## Logical Operators ##########################")
	maxMph := 50
	passengerCapacity := 4
	airbags := true

	familyCar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familyCar && !sportsCar

	fmt.Println(familyCar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

	// Converting, Parsing, and Formatting Values
	fmt.Println("########################## Converting, Parsing, and Formatting Values ##########################")
	kayak := 275
	soccerBall := 27.1

	// total := kayak + soccerBall
	total := float64(kayak) + soccerBall

	fmt.Println(total)

	/*
		Explicit conversions can be used only when the value can be represented in the target type. This means you can convert between numeric types and between strings and runes, but other combinations, such as converting int values to bool values, are not supported
	*/

	newTotal := kayak + int(soccerBall)
	fmt.Println(newTotal)
	fmt.Println(int8(newTotal)) //The int8 used in the second explicit conversion is too small to represent the int value 294 and so the variable overflows

	// math package functions
	fmt.Println("########################## Ceil, Floor, Round, RoundToEven ##########################")

	newTotal2 := kayak + int(math.Round(soccerBall))
	fmt.Println(newTotal2)

	ceilSoccerBall := int(math.Ceil(soccerBall))
	newTotal3 := kayak + ceilSoccerBall

	floorSoccerBall := int(math.Floor(soccerBall))
	newTotal4 := kayak + floorSoccerBall

	roundToEvenSoccerBall := int(math.RoundToEven(soccerBall))
	fmt.Println(roundToEvenSoccerBall % 2) // This is broken.
	newTotal5 := kayak + roundToEvenSoccerBall

	fmt.Println(ceilSoccerBall, floorSoccerBall, roundToEvenSoccerBall)
	fmt.Println(newTotal3, newTotal4, newTotal5)
}
