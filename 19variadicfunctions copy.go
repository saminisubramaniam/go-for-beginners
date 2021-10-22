package main

import (
	"fmt"
	"reflect"
)

func main() {

	welcome("Mohona")
	welcome("Oma", "Shira")
	welcome("Huma", "Leki", "Kira")
	welcome("Huma", "Leki", "Kira", "Moka")

	//calculate area for rectangle
	fmt.Println("Rectangle area:", calculate("rectangle", 100, 20))
	//calculate area for square
	fmt.Println("Square area: ", calculate("square", 25))

}

func welcome(names ...string) {
	fmt.Print(names, " is ")
	fmt.Print(reflect.ValueOf(names).Kind(), " with values: ")
	fmt.Println()
	//display the slice values in front of each output of line 24
	//use loops
}

func calculate(shape string, dimension ...int) (area int) {
	area = 1
	if shape == "square" {
		for _, v := range dimension {
			area = v * v
		}
	}

	if shape == "rectangle" {
		for _, v := range dimension {
			area *= v
			//area = area * v
		}
	}
	return
}

/*
variadic functions -> will accept varible
count of arguments of same datatype
*/
