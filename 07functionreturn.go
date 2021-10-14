package main

import "fmt"

func main() {
	fmt.Println(addition(4, 4))
	fmt.Println(multiplication(6, 6))
	var addResult = addition(5, 5)
	fmt.Println(addResult)
	fmt.Println(calulator("addition", 4, 3))
	fmt.Println(calulator("add", 9, 3))
	var out1, out2, unused, allowed = calulator("sum", 10, 8)
	fmt.Println(out1, out2, unused, allowed)
}

func addition(n1 int, n2 int) int {
	return n1 + n2
}

//return function variables as return values
func multiplication(n1 int, n2 int) (result int) {
	result = n1 * n2
	return
}

//return more than one function variables as return values
func calulator(calculation string, n1 int, n2 int) (c string, result int, status string, check bool) {
	c = calculation
	result = n1 + n2
	status = "alive"
	return
}
