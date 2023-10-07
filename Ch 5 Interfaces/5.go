/*
	Interfaces Quiz

	Q. Go uses the ____ keyword to show that a type implements an interface

	1. there is no keyword in go
	2. inherits
	3. implements
	4. fulfills

	A. 1
	Go에는 구현이라는 키워드가 따로 존재하지 않는다. 암시적으로 처리하기 때문
*/

package main

type shape interface {
	area() float64
}

type circle struct {
	//radius int
	radius float64
}

func (c *circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
