// Interfaces in go
/*
	구조체가 필드들의 집합체라면 interface는 메서드들의 집합체
	interface는 타입이 구현해야하는 메서드 원형(Prototype)을 정의한다.
	하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 interface가 갖는
	모든 메서드를 구현하면 된다.

	과제
	1. getMessage() message interface의 요구 사항으로 메서드 추가하기
	2. message interface 메서드를 통해 얻은 메시지 출력해 sendMessage 함수 완성
*/
package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("====================================\n")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}
