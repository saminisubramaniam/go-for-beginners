package main

import "fmt"

func main() {

	var counter = 25

	//if condition
	//check if value of counter is between 0 and 10, including 0 and 10
	//true && true .... => true
	//true && false ... => false
	//true || true .... => true
	//true || false ...  => true
	if counter <= 10 && counter >= 0 {
		fmt.Println("Counter less than or equal to 10  and greater than 0")
	}
	if counter >= 11 {
		fmt.Println("Counter more than 10")
	}

	//if else
	if counter <= 10 {
		//do something
		fmt.Println("If else -> Counter less than or equal to 10")
	} else {
		//do this thing
		fmt.Println("if else -> Counter more than 10")
	}

	//if else if else....

	if counter > 0 && counter < 10 {
		fmt.Println("Between 0 and 10")
	} else if counter > 10 && counter < 20 {
		fmt.Println("Between 10 and 20")
	} else if counter > 20 && counter < 30 {
		fmt.Println("Between 20 and 30")
	} else {
		fmt.Println("You missed it!!!!")
	}

	//declare and use the variable just after initialization
	var check int
	if check = 70; check > 70 {
		fmt.Println("Check pass!", check)
	} else {
		fmt.Printf("Check no pass as less than %v!", check)
	}

}
