/*
	Variadic
	가변성

	많은 함수, 특히 표준 라이브러리의 함수는 임의 개수의 최종 인수를 사용할 수 있다.
	이는 함수 서명에 "..." 구문을 사용하여 수행된다.
	(이게 뭔 소리인가... 코드로 직접 보면서 이해해보자)

	가변 함수는 가변 인수를 slice 받는다.
	func concat(strs ...string) string {
	    final := ""
	    // strs is just a slice of strings
	    for _, str := range strs {
	        final += str
	    }
	    return final
	}

	func main() {
	    final := concat("Hello ", "there ", "friend!")
	    fmt.Println(final)
	    // Output: Hello there friend!
	}

	약간.. python의 말 그대로 요소들을 추가하는...
	for i in itemList:
		final += i
	return final //을 보는 느낌


	----------------------------------


	스프레드 연산자

	func printStrings(strings ...string) {
		for i := 0; i < len(strings); i++ {
			fmt.Println(strings[i])
		}
	}

	func main() {
	    names := []string{"bob", "sue", "alice"}
	    printStrings(names...)
	}

*/

package main

import "fmt"

func sum(nums ...float64) float64 {
	// ?
	totalCost := 0.0
	for _, num := range nums {
		totalCost += num
	}
	return totalCost

}

// don't edit below this line

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}
