package main

import (
	"fmt"
	"reflect"
)

func main() {
	//copy array
	var scores = [4]int{88, 71, 11, 13}
	fmt.Println(scores)

	var scoresCopy = scores //copy by value
	fmt.Println(scoresCopy)
	scoresCopy[0] = 17
	fmt.Println(scoresCopy)

	//copy by reference
	var scoresReference = &scores //using & copy the memory location of scores
	fmt.Println(scoresReference)
	//888 is updated in both: scores and scoresReference
	scoresReference[0] = 888
	fmt.Println(*scoresReference)
	fmt.Println(scores)
	//update in scores
	scores[3] = 31
	fmt.Println(*scoresReference)
	fmt.Println(scores)

	//filter using :
	var friends = [8]string{"Oma", "Loki", "Paneer", "Tikka", "Alien", "Masala", "Dosa"}
	fmt.Println(friends)
	//print from 0 index to 1 index
	fmt.Println(friends[0:2])
	//same as above - 0 is ignored
	fmt.Println(friends[:2])
	//print from index 2 to 4
	fmt.Println(friends[2:5])
	//print from index 2 till last
	fmt.Println(friends[2:])
	//all values from array
	fmt.Println(friends)
	fmt.Println(friends[:])
	fmt.Println(friends[0:])
	fmt.Println(friends[0:len(friends)])
	//assume you do not know the size of array, then
	//how can you print last 2 elements of an array
	fmt.Println(friends[len(friends)-3:])
	//below works only for python
	//fmt.Println(friends[-2:0])

	fmt.Printf("%T \n", scores)
	fmt.Println(reflect.TypeOf(scores).Kind())
}
