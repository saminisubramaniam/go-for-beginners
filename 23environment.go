package main

import (
	"fmt"
	"os"
)

func main() {
	pathVariable := os.Getenv("NUMBER_OF_PROCESSORS")
	fmt.Println(pathVariable)

	//os.Setenv("TEST", "test")

	//what if environment variable is not set
	environmentVariable := "OS"
	value, isPresent := os.LookupEnv(environmentVariable)
	if isPresent {
		fmt.Printf("%v:%v", environmentVariable, value)
	} else {
		fmt.Printf("%v is not defined.", environmentVariable)
	}
}
