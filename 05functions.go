package main

import "fmt"

func main() {
	welcome()
	fmt.Println()
	learning()
}

func welcome() {
	var moon string = "Europa"
	fmt.Printf("I am %v, them moon of Jupiter!", moon)
}

func learning() {
	var moon = "Go"
	fmt.Printf("I am learning %v programming!", moon)
}
