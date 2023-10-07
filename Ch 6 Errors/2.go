/*
	Formatting Strings Review
	문자열 형식 검토

	Go에서 문자열 형식을 지정하는 편리한 방법은
	표준 라이브러리의 fmt.Sprintf 함수를 사용하는 것이다.

	이는 JavaScirpt의 내장 템플릿 리터럴과 유사한 문자열 보간 기능이다.
	하위 %v 문자열은 사용자가 원하는 유형의 기본 형식을 사용

	const name = "Kim"
	const age = 22
	s := fmt.Sprintf("%v is %v years old.", name, age)
	// s = "Kim is 22 years old."

	반올림 사례
	fmt.Printf("I am %f years old", 10.523)
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old", 10.523)
	// I am 10.52 years old

*/

/*
	ASSIGNMENT
	We need better error logs for our backend developers to help them debug their code.

	Complete the getSMSErrorString() function. It should return a string with this format:

	"SMS that costs $COST to be sent to 'RECIPIENT' can not be sent"

	COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
	RECIPIENT is the stringified representation of the recipient's phone number
*/

package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {

	return fmt.Sprintf("SMS that costs %.2f to sent to %v can not be sent", cost, recipient)

}

// don't edit below this line

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
}
