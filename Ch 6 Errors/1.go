/*
	The Error Interface

	Golang은 오류를 error 값으로 표현

	간단한 내장 오류 인터페이스
	type error interface {
    	Error() string
	}

	함수에서 문제가 발생할 수 있는 경우 해당 함수는 마지막 반환 값으로 error를 반환해야한다.
	error를 반환할 수 있는 함수를 호출하는 모든 코드는 오류가 있는지 테스트하여 오류를 처리해야한다.


*/

package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	// ?
	msgToCustom, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}

	msgToSpouse_, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return msgToSpouse_ + msgToCustom, nil
}

// don't edit below this line

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}
