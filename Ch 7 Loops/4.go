/*
	Fizzbuzz
	Go는 표준 연산자를 지원한다.

	나머지
	7 % 3 // 1

	Logical AND
	true && false // false
	true && true // true

	Logical OR
	true || false // true
	false || false // false

	피즈버즈 게임 만들기
*/

package main

import (
	"fmt"
)

func fizzbuzz() {

	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Print("fizz")
		} else if i%5 == 0 {
			fmt.Print("buzz")
		} else {
			fmt.Print(i)
		}

	}

}

// don't touch below this line

func main() {
	fizzbuzz()
}
