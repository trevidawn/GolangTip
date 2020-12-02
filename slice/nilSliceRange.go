package main

import "fmt"

func main() {
	var slice1 []int
	var slice2 *[]int

	for _, v := range slice1 {
		fmt.Println(v)
	}

	for _, v := range *slice2 {
		fmt.Println(v)
	}
	// error occur
	// panic: runtime error: invalid memory address or nil pointer dereference

}
