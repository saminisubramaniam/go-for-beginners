package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(reflect.ValueOf(currentTime).Kind())
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())
	fmt.Println(currentTime.Hour(), ":", currentTime.Minute(), ":", currentTime.Second())

	//counter will tick every 1 second
	counter := time.NewTicker(time.Second)
	for timer := range counter.C {
		currentTime = time.Now()
		fmt.Println(timer, " at ", currentTime.Hour(), ":", currentTime.Minute(), ":", currentTime.Second())
		//fmt.Println(currentTime.Hour(), ":", currentTime.Minute(), ":", currentTime.Second())
	}

}
