package main

import (
	"fmt"
	"math"
	"strconv"
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
	soccerBall := 19.5

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
	fmt.Println(roundToEvenSoccerBall % 2) // This is broken for 27.10
	newTotal5 := kayak + roundToEvenSoccerBall

	fmt.Println(ceilSoccerBall, floorSoccerBall, roundToEvenSoccerBall)
	fmt.Println(newTotal3, newTotal4, newTotal5)

	// parse from strings
	// The Go standard library includes the strconv package, which provides functions for converting string values to the other basic data types
	// ParseBool(str) ParseFloat(str, size) ParseInt(str, base, size) ParseUint(str, base, size) Atoi(str)
	val1 := "true"
	val2 := "false"
	val3 := "not true"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)

	if b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse:", val1)
	}

	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)

	val4 := "0"

	if bool4, b4err := strconv.ParseBool(val4); b4err == nil {
		fmt.Println("Parsed value:", bool4)
	} else {
		fmt.Println("Cannot parse", val4)
	}

	val5 := "100"

	int5, int5err := strconv.ParseInt(val5, 0, 8)

	if int5err == nil {
		smallInt := int8(int5)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val5)
	}

	val6 := "500"

	int6, int6err := strconv.ParseInt(val6, 0, 8)

	if int6err == nil {
		fmt.Println("Parsed value:", int6)
	} else {
		fmt.Println("Cannot parse", val6, int6err)
	}

	// Parsing Binary, Octal, and Hexadecimal Integers
	val7 := "100"
	int7, int7err := strconv.ParseInt(val7, 2, 8)

	if int7err == nil {
		smallInt := int8(int7)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val7, int7err)
	}

	val9 := "0b1100100" // need to skip val8 because of int8
	int9, int9err := strconv.ParseInt(val9, 0, 8)
	if int9err == nil {
		smallInt := int8(int9)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val9, int9err)
	}

	val10 := "100"
	int10, int10err := strconv.Atoi(val10)
	if int10err == nil {
		var intResult int = int10
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val10, int9err)
	}

	val11 := "48.95"
	float11, float11err := strconv.ParseFloat(val11, 64)
	if float11err == nil {
		fmt.Println("Parsed value:", float11)
	} else {
		fmt.Println("Cannot parse", val11, float11err)
	}

	val12 := "4.895e+01"
	float12, float12err := strconv.ParseFloat(val12, 64)
	if float11err == nil {
		fmt.Println("Parsed value:", float12)
	} else {
		fmt.Println("Cannot parse", val12, float12err)
	}

	val13 := true
	val14 := false

	str13 := strconv.FormatBool(val13)
	str14 := strconv.FormatBool(val14)

	fmt.Println("Formatted value 13: " + str13)
	fmt.Println("Formatted value 14: " + str14)

	val15 := 275
	// base10string := strconv.FormatInt(int64(val15), 10)
	base10string := strconv.Itoa(val15)
	base2string := strconv.FormatInt(int64(val15), 2)

	fmt.Println("Base 10: " + base10string)
	fmt.Println("Base 2: " + base2string)

	val16 := 49.95
	Fstring := strconv.FormatFloat(val16, 'f', 2, 64)
	Estring := strconv.FormatFloat(val16, 'e', -1, 64)
	Gstring := strconv.FormatFloat(val16, 'g', -1, 64)

	fmt.Println("Format F: " + Fstring)
	fmt.Println("Format E: " + Estring)
	fmt.Println("Format G: " + Gstring)

	// The special value -1 can be used to select the smallest number of digits that will create
	//  a string that can be parsed back into the same floating-point value without a loss of precision.
}
