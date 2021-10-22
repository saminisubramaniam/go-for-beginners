package main

import (
	"fmt"
	"reflect"
)

func main() {
	//map contains key value pair
	//key: value
	//var varname = map[datatype-of-key]datatype-of-value {"key1":"value1", "key2":"value2"}
	var student = map[string]int{"okie": 22, "yuko": 31, "mori": 54, "pico": 21, "luka": 31, "loki": 18}
	fmt.Println(student)
	fmt.Println(reflect.ValueOf(student).Kind())
	fmt.Println(student["okie"])
	//if the key does not exist then it returns 0
	fmt.Println(student["ok"])
	student["okie"] = 44
	fmt.Println(student["okie"])
	fmt.Println(student)
	//iterate over map
	for key, value := range student {
		fmt.Println("Key: ", key, "and", "Value: ", value)
	}

	//declaring map using make
	studentScore := make(map[string]int)
	fmt.Println(studentScore)
	studentScore["go"] = 20
	studentScore["html"] = 20
	fmt.Println(studentScore)
	studentScore["go"] = 22
	fmt.Println(studentScore)

	//delete value from map
	studentScore["css"] = 15
	studentScore["javascript"] = 22
	fmt.Println(studentScore)

	delete(studentScore, "html")
	fmt.Println(studentScore)
	//for the key which does not exist, the value is 0
	fmt.Println(studentScore["node"])
	fmt.Println(len(studentScore))

	//zero value map
	//newMap := make(map[string]string)
	var newMap map[string]string
	if newMap == nil {
		fmt.Println("newMap is nil")
	} else {
		fmt.Println("newMap is not nil")
	}

	//append all keys in an array
	var students []string
	fmt.Println(students)
	for key, _ := range student {
		students = append(students, key)
	}
	fmt.Println(students)

	//append all values in an array
	var scores []int
	fmt.Println(scores)
	for _, value := range student {
		scores = append(scores, value)
	}
	fmt.Println(scores)

}
