package main

import (
	"fmt"
	"os"
)

func main() {
	pointerToFile, error := os.Create("filename.txt")
	if error != nil {
		fmt.Println("1:", error)
		return
	}

	message := "Hello from Go. \nWriting to file.\nWorking fine."
	count, errorWriting := pointerToFile.WriteString(message)
	if errorWriting != nil {
		fmt.Println("2:", errorWriting)
		pointerToFile.Close()
		return
	}
	fmt.Println("Count of bytes written to file:", count)
	errorClose := pointerToFile.Close()
	if errorClose != nil {
		fmt.Println("3:", errorClose)
		return
	}
}
