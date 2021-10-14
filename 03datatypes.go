package main

import (
	"fmt"
)

func main() {

	//int datatype
	var score int = 88
	//boolean datatype
	var check bool
	//string datatype
	var name = "AAA"
	//declare multiple variables
	var city, location, country, population = "KL", "4th Street", "Malayasia", 3133
	fmt.Println(location, city, country, "with population of", population)
	//employee - declare related variables in block
	var (
		employeeName  = "Oki"
		employeeAge   = 31
		employeeScore = 88
	)

	fmt.Println(employeeName, employeeAge, employeeScore)

	fmt.Println(score)
	fmt.Println(check)

	fmt.Printf("%T, %T, %T \n", score, check, name)
	fmt.Printf("%v, %7v, %v \n", score, check, name)

	fmt.Printf("My score is %v and I am from %v. \n", score, name)
	fmt.Println("My score is", score, "and I am from", name, ".")
	//below will not work as "My score is" is string and score is integer
	//fmt.Println("My score is" + score, "and I am from", name, ".")
}
