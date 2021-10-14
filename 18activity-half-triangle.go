package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Input the rows for half triangle:")
	var rowsCount int
	fmt.Scanln(&rowsCount)
	fmt.Println("Print half traingle for", rowsCount)
	for i := 1; i <= rowsCount; i++ {
		// for i := 1; i <= 6; i++ {
		fmt.Printf("%s\n", strings.Repeat("$", i))
	}
}

/*
$
$$
$$$
$$$$
$$$$$
$$$$$$


*/
