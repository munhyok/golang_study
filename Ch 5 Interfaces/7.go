/*
	Multiple Interfaces

	type은 Go에서 원하는 수의 interface를 구현할 수 있다.

	예를 들어 빈(empty) interface인 interface{}는 요구사항이 없기 때문에
	항상 모든 타입에 의해 구현된다.


	ASSIGNMENT
	Complete the required methods so that the email type implements both the expense and printer interfaces.

	COST()
	If the email is not "subscribed", then the cost is 0.05 for each character in the body.
	If it is, then the cost is 0.01 per character.

	Return the total cost of the entire email.

	PRINT()
	The print method should print to standard out the email's body text.
*/

package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if !e.isSubscribed { //구독 유무 확인
		return float64(len(e.body)) * 0.5 //body length type는 int이기 때문에 float64로 형변환
	}
	return float64(len(e.body)) * 0.1

}

func (e email) print() {
	fmt.Println(e.body)

}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================\n")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}
