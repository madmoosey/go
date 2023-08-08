package main

import (
	"fmt"
	"strconv"
)

func main() {
	kayakPrice := 275.00
	fmt.Println("Price:", kayakPrice)

	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}

	/*
		Go is less flexible than other languages when it comes to the syntax for if statements and other flow control statements.
		First, the braces cannot be omitted even when there is only one statement in the code block.
		Second, the opening brace must appear on the same line as the flow control keyword and cannot appear on the following line.
		third, if you want to split a long expression onto multiple lines, you cannot break the line after a value or variable name.
		Each clause in an if statement has its own scope, which means that variables can be accessed only within the clause in which they are defined. It also means you can use the same variable name for different purposes in separate clauses
	*/

	/*
		Go allows an if statement to use an initialization statement, which is executed before the if statement’s expression is evaluated.
		The initialization statement is restricted to a Go simple statement, which means—in broad terms—that the statement can define a new variable, assign a new value to an existing variable, or invoke a function
	*/

	priceString := "275"

	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}

	// For loops

	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	newCounter := 0
	for newCounter <= 3 {
		fmt.Println("New Counter:", newCounter)
		newCounter++
	}

	for conCounter := 0; conCounter <= 3; conCounter++ {
		if conCounter == 1 {
			continue
		}
		fmt.Println("Count Counter:", conCounter)
	}

	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}

	for onlyIndex := range product {
		fmt.Println("Index:", onlyIndex)
	}

	for _, onlyCharacter := range product {
		fmt.Println("Character:", string(onlyCharacter))
	}

	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}

	for productIndex, productElement := range products {
		fmt.Println("Index:", productIndex, "Element:", productElement)
	}

	for switchIndex, switchCharacter := range product {
		switch switchCharacter {
		case 'K', 'k':
			if switchCharacter == 'k' {
				fmt.Println("Lowercase k at position", switchIndex)
				break
			}
			fmt.Println("Uppercase K at position", switchIndex)
		case 'y':
			fmt.Println("y at position", switchIndex)
		}

	}

	// Fallthrough

	for switchIndex, switchCharacter := range product {
		switch switchCharacter {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough

		// The first case statement contains the fallthrough keyword, which means that execution will continue to the statements in the next case statement
		case 'k':
			fmt.Println("k at position", switchIndex)
		case 'y':
			fmt.Println("y at position", switchIndex)
		default:
			fmt.Println("Character", string(switchCharacter), "at position", switchIndex)
		}
	}

	// A switch statement can be defined with an initialization statement, which can be a helpful way of preparing the comparison value so that it can be referenced within case statement

	for initCounter := 0; initCounter < 20; initCounter++ {
		switch val := initCounter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}

	for noCompCounter := 0; noCompCounter < 10; noCompCounter++ {
		switch {
		case noCompCounter == 0:
			fmt.Println("Zero value")
		case noCompCounter < 3:
			fmt.Println(noCompCounter, "is < 3")
		case noCompCounter >= 3 && noCompCounter < 7:
			fmt.Println(noCompCounter, "is >= 3 && < 7")
		default:
			fmt.Println(noCompCounter, "is >= 7")
		}
	}

	// Label statements allow execution to jump to a different point, giving greater flexibility than other flow control features

	labelCounter := 0
target:
	fmt.Println("Label counter", labelCounter)
	labelCounter++
	if labelCounter < 5 {
		goto target
	}
}
