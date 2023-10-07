/*
	The Error Interface

	오류는 인터페이스일 뿐이므로 인터페이스를 구현하는 사용자 정의 유형을 직접 만들수 있다.

	예시
	type userError struct {
	    name string
	}

	func (e userError) Error() string {
	    return fmt.Sprintf("%v has a problem with their account", e.name)
	}

	func sendSMS(msg, userName string) error {
	    if !canSendToUser(userName) {
	        return userError{name: userName}
	    }
	    ...
	}

*/

package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", e.dividend)
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convert the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}
