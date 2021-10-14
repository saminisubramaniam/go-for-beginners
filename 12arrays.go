package main

import "fmt"

func main() {
	//declare array
	var scores = [4]int{88, 71, 11, 13}
	fmt.Println(scores)

	//... will automatically get the size of the array
	var scoresAuto = [...]int{88, 71, 11, 13}
	fmt.Println(scoresAuto)

	var friends = [6]string{"Oma", "Loki", "Pani", "Tika", "Alien"}
	fmt.Println(friends)
	//use index positon to access individual elements in an array
	//index always starts from 0
	fmt.Println(scores[0])
	fmt.Println(friends[1])
	friends[3] = "Bindi"
	fmt.Println(friends)

	fmt.Println("Length of array scores: ", len(scores))
	fmt.Println("Length of array friends: ", len(friends))

	//empty array
	var empty = [8]int{}
	fmt.Println("Length of array empty: ", len(empty))
	fmt.Println(empty)

	//intialize the values in array at specific index position
	var specific = [9]int{4: 8, 8: 17}
	fmt.Println(specific)

	//2d array
	var matrix = [3][4]int{
		{0, 1, 2, 3},   //row with index position 0
		{4, 5, 6, 7},   //row with index position 1
		{8, 9, 10, 11}} //row with index position 2

	fmt.Println(matrix)
	//print row 1
	fmt.Println(matrix[1])
	//print 7
	fmt.Println(matrix[1][3])

}
