/*
	Append
	당연하게도 Go도 당연히 Append를 지원한다

	func append(slice []Type, elems ...Type) []Type

	전부 유효한 코드
	slice = append(slice, oneThing)
	slice = append(slice, firstThing, secondThing)
	slice = append(slice, anotherSlice...)


	https://hwan-shell.tistory.com/346
	정리한 링크

*/

package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	// ?
	costsByDay := []float64{} //가변 배열 선언
	for i := 0; i < len(costs); i++ {
		cost := costs[i] // costs[0] = {0, 1.0}

		for cost.day >= len(costsByDay) { // 0 >= 0 -> true
			costsByDay = append(costsByDay, 0.0) //costByDay에 0.0 추가
		}

		costsByDay[cost.day] += cost.value

	}
	return costsByDay

}

// dont edit below this line

func test(costs []cost) {
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByDay := getCostsByDay(costs)
	fmt.Println("Costs by day:")
	for i := 0; i < len(costsByDay); i++ {
		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
	test([]cost{
		{0, 1.0},
		{10, 2.0},
		{3, 3.1},
		{2, 2.5},
		{1, 3.6},
		{2, 2.7},
		{4, 56.34},
		{13, 2.34},
		{28, 1.34},
		{25, 2.34},
		{30, 4.34},
	})
}
