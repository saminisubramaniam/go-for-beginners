package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("--------------------")
	counter := 1
	for counter < 4 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println("--------------------")
	for i := 0; ; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
	//range keyword
	friend := "Oki"
	for range friend {
		fmt.Println("Hello....")
	}
	scores := [4]int{88, 44, 33, 22}
	for i := range scores {
		fmt.Println("Value at index", i, "is", scores[i])
	}
	fmt.Println("--------------------")
	for i, value := range scores {
		fmt.Println("Value at index", i, "is", value)
	}
	fmt.Println("--------------------")
	//iterating on map using range
	keyvalue := map[string]string{
		"id":        "1",
		"name":      "AAA",
		"location":  "KL",
		"country":   "Malaysia",
		"continent": "Asia"}

	for key, value := range keyvalue {
		fmt.Println(key, ":", value)
	}

}
