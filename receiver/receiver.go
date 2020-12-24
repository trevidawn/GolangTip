package main

import "fmt"

type A struct {
	i int
	s string
}

func (a A) ValueFunc() {
	fmt.Printf("%+v\n", a)
}

func (a *A) PointerFunc() {
	fmt.Printf("%+v\n", a)
}

func main() {

	var (
		A_Value   A
		A_Pointer *A
	)

	A_Value = A{}
	A_Pointer = &A{}

	A_Value.ValueFunc()
	A_Value.PointerFunc()
	A_Pointer.ValueFunc()
	A_Pointer.PointerFunc()

	// 의외로 Value, Pointer Type에 대해서 Value Receiver, Pointer Receiver 모두 호출이 가능하다.
}
