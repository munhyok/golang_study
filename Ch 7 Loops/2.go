/*
Omitting Conditions from a for loop in Go

for loop의 section을 생략할 수 있는데

기존 구조가

	for INITIAL; CONDITION; AFTER{
	  // do something
	}

라면

CONDITION 부분을 생력할 수 있다.

	for INITIAL; ; AFTER {
	  // do something forever
	}

이게 사실상의 While문이 아닐까 싶다 신기하네
*/
package main

import (
	"fmt"
)

func maxMessages(thresh float64) int {
	cost := 1.0
	for i := 0; ; i++ {

		cost += float64(i) * 0.01
		if cost > thresh {
			return i
		}
	}

}

// don't edit below this line

func test(thresh float64) {
	fmt.Printf("Threshold: %.2f\n", thresh)
	max := maxMessages(thresh)
	fmt.Printf("Maximum messages that can be sent: = %v\n", max)
	fmt.Println("===============================================================")
}

func main() {
	test(10.00)
	test(20.00)
	test(30.00)
	test(5.1)
	test(40.00)
	test(50.00)
}
