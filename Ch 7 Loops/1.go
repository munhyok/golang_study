/*
	Loops in Go
	Go의 기본 루프는 C와 비슷하다

	for i := 0; i < 10; i++ {
		fmt.Println("요런 느낌적인 느낌")
	}


*/

package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	cost := 1.0
	for i := 0; i < numMessages; i++ {

		cost += float64(i) * 0.01

	}
	return cost
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
