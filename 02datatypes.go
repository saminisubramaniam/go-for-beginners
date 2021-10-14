package main

import "fmt"

func main() {
	//var name datatype
	/*
		i am
		the
		multiline comment
	*/
	var name string = "Airasia Academy"
	fmt.Println(name)

	var noDataTypeVaraible = "Alien"
	fmt.Println(noDataTypeVaraible)

	fmt.Print("Hello without new line.")
	fmt.Print("I am in same line.")
	fmt.Println("")
	//using verbs to format the output
	fmt.Printf("I am from %v and I am an %v. \n", name, noDataTypeVaraible)
	//using space with verbs
	fmt.Printf("I am another %10 and I am at %-30v. \n", noDataTypeVaraible, name)
	//get the datatype of the varivable
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", noDataTypeVaraible)

}
