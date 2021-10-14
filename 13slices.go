package main

import (
	"fmt"
	"reflect"
)

func main() {
	var dynamicArray = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(dynamicArray)
	fmt.Println(len(dynamicArray))
	dynamicArray[0] = 0
	fmt.Println(dynamicArray)
	//dynamicArray[10] = 11
	fmt.Println(dynamicArray)

	var globalfriends = []string{"obb", "owio", "chk"}
	fmt.Println(globalfriends)

	fmt.Println(reflect.TypeOf(dynamicArray).Kind())

	//get better control in declaring slice
	//make(datatype, lenght, capacity)
	var mySlice = make([]int, 2, 5)
	fmt.Println(mySlice)
	mySlice[1] = 10
	fmt.Println(mySlice)
	fmt.Println("Capacity of mySlice:", cap(mySlice))
	fmt.Println("Length of mySlice:", len(mySlice))

	//declare slice with new
	//50 is capacity
	var newSlice1 = new([50]int)
	fmt.Println(newSlice1)

	//declare slice with new
	//50 is capacity
	//lenght is 5
	var newSlice2 = new([50]int)[0:5]
	fmt.Println(newSlice2)
	fmt.Println("Capacity of newSlice2:", cap(newSlice2))
	fmt.Println("Length of newSlice2:", len(newSlice2))

	newSlice2 = new([50]int)[0:10]
	fmt.Println("Capacity of newSlice2:", cap(newSlice2))
	fmt.Println("New Length of newSlice2:", len(newSlice2))

	newSlice2 = new([50]int)[0:50]
	fmt.Println("Capacity of newSlice2:", cap(newSlice2))
	fmt.Println("New Length of newSlice2:", len(newSlice2))

}
