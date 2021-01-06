package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	var i int
	for range ticker.C {
		if i == 3 {
			i++
			continue
		}
		fmt.Printf("tick %d!\n", i)
		i++
	}
}