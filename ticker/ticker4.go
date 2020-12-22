package main

import (
	"fmt"
	"time"
)

/*
	간단한 select로 ticker 테스트
*/

func main() {

	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(time.Second * 2)
	ticker3 := time.NewTicker(time.Second * 3)

	for {
		select {
		case <-ticker1.C:
			fmt.Printf("%s : %v\n", "ticker1", time.Now())
		case <-ticker2.C:
			fmt.Printf("%s : %v\n", "ticker2", time.Now())
		case <-ticker3.C:
			fmt.Printf("%s : %v\n", "ticker3", time.Now())
		}
	}
}
