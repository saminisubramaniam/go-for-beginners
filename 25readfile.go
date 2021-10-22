package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filedata, error := ioutil.ReadFile("filename.txt")
	if error != nil {
		fmt.Println("File read error", error)
		return
	}

	fmt.Println("Reading file was success....")
	fmt.Println(filedata)
	fmt.Println(string(filedata))
}
