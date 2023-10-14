/*
	Make

	대부분의 경우 Slice의 기본 배열에 대해 생각할 필요가 없다.
	make라는 함수를 이용해 slice를 만들 수 있다.

	func make([]T, len, cap) []T
	[]T : type 변수 타입 지정
	len : lenght 배열 길이 결정
	cap : capacity 용량 보통 length 길이에 맞춰진다.
	그래서 cap를 생략할 수 있다. 알아서 len에 맞춰서 늘어나기 때문



	mySlice := make([]int, 5, 10)
	the capacity argument is usually omitted and defaults to the length
	mySlice := make([]int, 5)

	len 함수 index 길이 구하는 함수
	mySlice := []string{"I", "love", "go"}
	fmt.Println(len(mySlice)) // 3

	cap 함수
	mySlice := []string{"I", "love", "go"}
	fmt.Println(cap(mySlice)) // 3


	mySlice := make([]int, 5, 10)
	fmt.Println(cap(mySlice)) // 10


	그럼 cap은 왜 있을까?
	메모리 사용량을 제한하기 위해 사용하는 용도

	배열하고 무슨 차이가 있지?
	-> 배열도 결국 길이를 정해 메모리 사용량을 정할 수 있지 않냐!

	예상하기엔 make를 사용하면 정적/동적 메모리로 나뉠 것 같다.

	예를들어
	var myInt = [10]int 라고 선언하면
	해당 배열의 크기는 고정이 된다. 값을 10개 밖에 넣지 못한다.

	하지만 make를 사용한다면
	mySlice := make([]int, 5, 10)
	처음 메모리 선언 시 len을 5로 설정했기 때문에
	초기 선언시 {0,0,0,0,0}로 선언은 되지만
	그 이후에 어떤 값을 더 넣어야한다면 무려 5개나 더 넣을 수 있다.
	cap은 말 그대로 해당 배열에 최대 몇 개의 수를 동적으로 넣을 수 있냐로 생각하면 될듯하다

	한 방에 최대 10명을 수용할 수 있다라고 이해하면 편할듯...?

	make함수에 cap을 설정해두지 않았다면 동적으로 len에 맞춰 알아서 늘어나고 줄어들 것이다.

	cap를 잘 이해하고 사용하면 메모리 관리에 도움이 될지도 모르겠다.

*/

package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	// ?
	messagesCost := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := 0.01 * float64(len(messages[i]))
		messagesCost[i] = cost
	}

	return messagesCost
}

// don't edit below this line

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}
