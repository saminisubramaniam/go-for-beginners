package main

import "fmt"

func main() {
	fmt.Println(factorial(100))
}

func factorial(n int) (output int) {
	//n = n  * (n-1) * (n-2) * (n-3) .... * 1
	//in recursive function, it is mandatory to use stop condition
	//to prevent recursion from going into infinite loop
	if n > 0 {
		output = n * factorial(n-1)
	} else {
		output = 1
	}

	return
}
