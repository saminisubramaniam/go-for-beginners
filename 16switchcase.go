package main

import "fmt"

func main() {

	var check = 22
	//basic switch case
	switch check {
	case 10:
		fmt.Println("Ten")
	case 11:
		fmt.Println("Eleven")
	case 12:
		fmt.Println("Twelve")
	default:
		fmt.Println("Something else")
	}

	//more than one case
	switch check {
	case 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
		fmt.Println("Tens")
	case 20, 21, 22, 23, 24, 25, 26, 27, 28, 29:
		fmt.Println("Twenties")
	case 30, 31, 32, 33, 34, 35, 36, 37, 38, 39:
		fmt.Println("Thirties")
	default:
		fmt.Println("Something else")
	}

	//more than one case
	switch {
	case check >= 10 && check < 20:
		fmt.Println("Tens")
	case check >= 20 && check < 30:
		fmt.Println("Twenties")
	case check >= 30 && check < 40:
		fmt.Println("Thirties")
	default:
		fmt.Println("Something else")
	}

}
