package main

import (
	"fmt"
	"time"
)

/*
	Stop한 Ticker를 Reset하면 Tick이 전송되는지 테스트
 */

func main() {
	ticker := time.NewTicker(time.Second)

	ticker.Stop()

	ticker.Reset(time.Second * 2)

	for {
		fmt.Println(<-ticker.C)
	}

	// 잘 전송된다.
	// Stop해도 실제 Channel을 닫진 않기 때문에, 채널 그대로 쓰는듯?
}
