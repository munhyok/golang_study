// Interface Implementation 인터페이스 구현하기
/*
	인터페이스는 암시적으로 구현된다.
	아래 유형은 주어진 인터페이스를 구현한다고 선언하지 않는다.
	인터페이스가 존재하고 유형에 적절한 메소드가 정의되어 있으면 해당 유형이 자동으로
	해당 인터페이스를 수행한다.

	과제
	employee interface는 다양한 직원 유형을 더 쉽게 처리할 수 있도록 보다
	일반적인 인터페이스를 만들어야한다.

	1. employee interface를 충족하도록 누락된 getSalary 메서드를 contractor 타입에 추가하기
	2. contractor의 급여는 시간당 급여에 연간 몇 시간을 곱한 것

*/

package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// ?

// don't touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func main() {
	test(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
}
