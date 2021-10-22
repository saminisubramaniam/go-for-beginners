package main

import "fmt"

type student struct {
	name     string
	score    int
	subject  string
	promoted bool
	presence int
}

func main() {
	var student1 = student{}
	fmt.Println(student1)

	var student2 = student{name: "Okie", score: 22, subject: "Computer Science"}
	fmt.Println(student2)

	student1.name = "Moki"
	student1.score = 44
	student1.subject = "Maths"
	fmt.Println(student1)

	var student3 = student{"Okie", 22, "Science", true, 100}
	fmt.Println(student3)

	fmt.Printf("%v has scored %v in %v with %v percent presence.\n", student3.name, student3.score, student3.subject, student3.presence)

	//declare instance using new keyword
	student4 := new(student)
	student4.name = "Luka"
	fmt.Println(student4)
}

/*
type structName struct{
	variable1 datatype1
	variable2 datatype2
	variable3 datatype3
}

customVarible := structName{}


*/
