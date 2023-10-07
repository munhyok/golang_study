//Ignore Return Values

package main

import "fmt"

func main() {
	firstName, _ := getNames() //언더바로 받아오는 값을 패스할 수 있다..

	fmt.Println("Welcome to Textio, ", firstName)

}

func getNames() (string, string) {
	return "Munhyeok", "Kang"
}
