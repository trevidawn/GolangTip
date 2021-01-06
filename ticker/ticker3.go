package main

import (
	"fmt"
	"time"
)

/*
	default ticker, nil ticker에 Stop, Reset, 채널 읽기를 수행했을 때, 어떤 결과가 나올것인가.
	각 라인을 주석처리하며 하나씩 테스트해봄
 */

var (
	defaultTicker time.Ticker
	nilTicker *time.Ticker
)

func main() {
	/*
	1. ticker Stop
	*/
	defaultTicker.Stop()

	//아무 에러 없음, nil ticker를 Stop하는 것은 문제가 되지 않는다.
	nilTicker.Stop()
	//segmentation err 발생! : invalid memory address or nil pointer dereference

	/*
	2. ticker.Reset
	*/
	defaultTicker.Reset(time.Second)
	/*
		의외로 panic 발생, Reset과정에 ticker 초기화과정이 포함될거라 생각했는데 아닌가봄
		Reset called on uninitialized Ticker
		추측하자면 ticker가 nil이기 때문에 ticker의 Channel이 없기 때문에 panic이 발생한거 같다.
	*/
	nilTicker.Reset(time.Second)

	//segmentation err 발생! : invalid memory address or nil pointer dereference


	/*
	3. 채널 읽기
	*/

	fmt.Println(<-defaultTicker.C)
	//에러 발생 : fatal error: all goroutines are asleep - deadlock!
	//nil Ticker의 Chan

	fmt.Println(<-nilTicker.C)
	//segmentation err 발생! : invalid memory address or nil pointer dereference
}
