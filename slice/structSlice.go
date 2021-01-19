package main

import (
	"fmt"
	"strconv"
)

/*
	struct slice는 변경될까?
	당연히 변경됨;; 다만 함수의 인자로 전달해서 해당함수에서 변경하면 안되지
*/

type SliceValue struct {
	i int
	s string
}

func main() {

	slice := make([]SliceValue, 0)

	for i := 0; i < 5; i++ {
		slice = append(slice, SliceValue{
			i: i,
			s: strconv.Itoa(i),
		})
	}

	slice[3].i = 100
	slice[3].s = "100"

	for idx, v := range slice {
		fmt.Printf("%d : %+v\n", idx, v)
	}
}
