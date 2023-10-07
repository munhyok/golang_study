/*
	Type Switches

	타입 스위치를 사용하면 일련의 여러 Type Assertion을 쉽게 수행할 수 있다.
	타입 스위치는 일반 스위치 문과 비슷하지만 사례에서는 값 대신 타입을 지정한다.

	예제

	func printNumericValue(num interface{}) {
		switch v := num.(type) {
		case int:
			fmt.Printf("%T\n", v) //%T = type 출력
		case string:
			fmt.Printf("%T\n", v)
		default:
			fmt.Printf("%T\n", v)
		}
	}

	func main() {
		printNumericValue(1)
		// prints "int"

		printNumericValue("1")
		// prints "string"

		printNumericValue(struct{}{})
		// prints "struct {}"
	}

*/

package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {

	switch v := e.(type) { //타입을 v 변수에 할당 expense -> cost() -> email,sms,invalid
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()

	default:
		return "", 0.0
	}

}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Wanna catch up later?",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "I don't need this",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}
