// The Benefits of named Returns 10~11

package main

import (
	"errors"
)

// 문서화를 할 땐 return 키워드에 반환 할 변수를 명시하자 + 협업할 때
func calculator_return(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

// naked return = return 키워드에 굳이 반환 할 변수를 명시하지 않아도 됨을 뜻한다.
// 짧은 코드의 함수는 naked return을 사용해도 쉽게 볼 수 있어서 함수 코드가 짧을 때만 사용하자
func calculator_naked(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul := a * b
	div := a / b
	return mul, div, nil
}
