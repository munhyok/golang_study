/*
	Type Assertions in GO

	Go에서 interface로 작업할 때 가끔 interface 값의 기본 유형에 액세스를 해야하는 경우가 있다.
	이럴 때 type assertion을 사용하여 인터페이스를 기본 유형으로 캐스팅할 수 있다.

	예제 코드
	type shape interface {
		area() float64
	}

	type circle struct {
		radius float64
	}

	// "c" is a new circle cast from "s"
	// which is an instance of a shape.
	// "ok" is a bool that is true if s was a circle
	// or false if s isn't a circle
	c, ok := s.(circle)
	if !ok {
		// s wasn't a circle
		log.Fatal("s is not a circle")
	}

	radius := c.radius

	솔직히 까고 말해서 예제로 봤을 때 뭔 소리인가 싶어서 직접 문제를 보면서 이해해보자
	생략된 부분이 많은 느낌이라..

	***나중에 다시 복습하기***

	ASSIGNMENT
	Implement the getExpenseReport function.

	If the expense is an email then it should return the email's toAddress and the cost of the email.
	If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
	If the expense has any other underlying type, just return an empty string and 0.0 for the cost.
*/

package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	// ?
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	sm, ok := e.(sms)

	if ok {
		return sm.toPhoneNumber, em.cost()
	}
	return "", 0.0
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
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}
