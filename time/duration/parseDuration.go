package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	durationString := ""
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(duration)
	// OUTPUT:
	// invalid duration ""
	// 빈 문자열은 파싱되지 않음
}
