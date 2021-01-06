package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	var i int
	for range ticker.C{
		fmt.Printf("%+v : %s\n", len(ticker.C), time.Now())
		if i == 3 {
			time.Sleep(time.Second * 5)
		}
		i++
	}
}