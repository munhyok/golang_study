/*
	Range

	파이썬 처럼 for문으로 요소를 넣는 방법이 있다.

	for INDEX, ELEMENT := range SLICE {

	}

	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
	    fmt.Println(i, fruit)
	}
	// 0 apple
	// 1 banana
	// 2 grape


*/

package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {

	for i, message := range msg {
		for _, badWord := range badWords {
			if message == badWord {
				return i

			}
		}
	}

	return -1

}

// don't touch below this line

func test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func main() {
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test(message, badWords)
}
