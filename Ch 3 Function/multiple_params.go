package main

import "fmt"

// 여러개의 파라미터를 묶어서 지정 가능 일일이 하나씩 지정할 필요가 없다.
func addToDatabase(hp, damage int, name string, level int) {
	fmt.Println(name, "\n레벨:", level, "\n체력:", hp, "\n데미지:", damage)
}

func main() {
	addToDatabase(10, 10, "kangmunhyeok", 30)
}
