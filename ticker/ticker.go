package main

import "time"

/*
	type Ticker struct {
		C <- chen Time // Ticker의 channel C를 통해 Tick이 일정 주기로 전송된다.
	}

	func NewTicker(d time.Duration) *Ticker
	func (t *Ticker) Reset(d Duration) // go 1.15 버전부터 사용이 가능하다.
	func (t *Ticker) Stop()

 */

func main() {
	ticker := time.NewTicker(time.Second)
	/*
		1초 주기의 신호를 보내는 Ticker를 생성한다.
		주기는 0보다 커야하며, 그렇지 않으면 panic이 발생하니 조심하자.
	 */

	ticker.Reset(time.Second * 2)
	/*
		ticker를 멈추고, 인자로 주어진 duration으로 새로운 주기를 설정하고 ticker를 시작한다.
		위 예시에서는 ticker를 멈추고 1초 주기를 2초로 변경
		2초 뒤에 tick이 전송됨
	 */

	ticker.Stop()
	/*
		ticker를 정지한다. 즉 tick을 channel C에 전송하지 않는다.
		실제 channel을 닫지는(close) 않는다.
	 */
}
