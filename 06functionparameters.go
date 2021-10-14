package main

import "fmt"

func main() {
	printTeamMember("Prafful Daga", 17)
	fmt.Println()
	printTeamMember("Oki San", 8)
}

func printTeamMember(name string, experience int) {
	fmt.Printf("%v has experience of %v years.", name, experience)
}
