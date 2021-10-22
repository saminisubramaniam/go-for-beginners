package main

import "fmt"

type student struct {
	name          string
	score         int
	subject       string
	studentSchool school
}

type school struct {
	schoolname string
	city       string
}

func main() {
	school1 := school{schoolname: "AAA", city: "KL"}
	student1 := student{
		name:          "Dori",
		score:         71,
		subject:       "Game Development",
		studentSchool: school1}

	fmt.Println(student1)

	school2 := new(school)
	student2 := new(student)

	school2.city = "LK"
	school2.schoolname = "RBA"
	student2.name = "Luma"
	student2.score = 44
	student2.subject = "Graphics"
	student2.studentSchool = *school2
	fmt.Println(student2)

	fmt.Printf("%v is from %v in %v", student2.name, student2.studentSchool.schoolname, student2.studentSchool.city)

}
