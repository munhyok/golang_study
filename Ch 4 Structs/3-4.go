// go anonymous structs

/*
	익명 구조체는 일반 구조체와 같지만 이름 없이 정의되므로 코드의 다른 곳에서
	참조할 수 없다.
*/

package main



//구조체를 만들고 즉시 인스턴스를 인스턴스화 시키는 예시
myCar := struct {
	Make 	string
	Model	string

} {
	Make: "tesla",
	Model: "model 3",
}

//익명 구조체를 다른 구조체 내의 필드로 중첩도 가능
type car struct {
	Make	string
	Model	string
	Height	int
	Width	int

	Wheel struct {
		Radius		int
		Material	string
	}
}

/*
	익명 구조체는 언제 사용해야할 까?
	
	일반적으론 Named 구조체를 사용하고 Named 구조체를 사용하면
	코드를 더 쉽게 읽고 이해할 수 있으며 재사용이 가능하다

	구조체를 다시 사용할 필요가 없다는 것을 알아도 가끔 익명구조체를 이용하는 경우도 존재한다
	예를 들어 HTTP 처리에서 일부 JSON 데이터 모양을 만드는데 사용한다.
	가만보면 JSON형식과 유사한 것 같기도...

	하지만 보통 한 번만 사용되도록 선언해 개발자가 실수로 다시 사용하려는 것을 방지할 수 있게
	선언하는 것이 합리적이다.
	

*/