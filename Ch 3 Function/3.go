// Declaration Syntax 3~5

package main

import "fmt"

func main() {
	var x int = 10 //일반 변수 선언
	var p *int     //poiner
	//var a [3]int array

	p = &x //p에 x의 메모리 주소 할당 p = x메모리 주소, *p = 10
	*p = 20
	/*하지만 20을 다시 선언하면서 x 메모리 주소를 참조하는 포인터p는 20으로 바뀜
	이때 x의 메모리 주소를 참조했기때문에 x값도 20으로 변경*/
	fmt.Println("p포인터\n")
	fmt.Println(&x) //x의 메모리 주소
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(x)

}
