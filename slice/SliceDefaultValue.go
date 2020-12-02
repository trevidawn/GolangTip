package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := []int{}
	var slice3 []int = nil
	var slice4 *[]int
	slice5 := &[]int{}

	fmt.Println(slice1) // []
	fmt.Println(slice2) // []
	fmt.Println(slice3) // []
	fmt.Println(slice4) // <nil>
	fmt.Println(slice5) // &[]

	if slice1 == nil {
		fmt.Println("slice1 is nil")
	}

	if slice2 == nil {
		fmt.Println("slice2 is nil")
	}

	if slice3 == nil {
		fmt.Println("slice3 is nil")
	}

	if slice4 == nil {
		fmt.Println("slice4 is nil")
	}

	if slice5 == nil {
		fmt.Println("slice5 is nil")
	}

	// output:
	// slice 1 is nil
	// slice 3 is nil
	// slice 4 is nil
}
