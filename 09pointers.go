package main

import "fmt"

func main() {
	var fname = "PD"
	//normal function call -> pass by value
	changeName(fname)
	//this will print fname declared in line 6.
	//this will not print fname updated in changeName function
	fmt.Println(fname)
	//use pointers -> pass memory address of fname
	changeNameNow(&fname)
	//this will print fname updated in changeNameNow function
	fmt.Println(fname)

}

func changeName(fn string) (fname string) {
	fname = "Mr. " + fn
	return
}

func changeNameNow(fn *string) {
	fmt.Println(fn)
	*fn = "OBB"
	//fname = "Mr. " + fn
	//return
}
