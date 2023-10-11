/*
	Continue
	루프의 현재 반복을 중지하고 다음 반복을 계속

	Break
	루프의 반복을 중지하고 루프 종료

	for i := 0; i < 10; i++ {
	  if i % 2 == 0 {
	    continue
	  }
	  fmt.Println(i)
	}
	// 1
	// 3
	// 5
	// 7
	// 9

	for i := 0; i < 10; i++ {
	  if i == 5 {
	    break
	  }
	  fmt.Println(i)
	}
	// 0
	// 1
	// 2
	// 3
	// 4


*/

package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}

		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(n)
		}

	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
