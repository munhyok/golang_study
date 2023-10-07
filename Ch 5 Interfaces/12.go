package main

import (
	"io"
	"os"
)

/*
	Clean Interfaces 12 ~ 15

	인터페이스를 깨끗하게 유지하기 위한 법칙
*/

/*
	1. 인터페이스를 작게 유지하기
	인터페이스는 아이디어나 개념을 정확하게 표현하는데 필요한 최소한의 동작을 정의하기 위한 것이다.

	다음은 최소한의 동작을 정의하기 좋은 예시인 표준 HTTP 패키지 예시
*/

type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

/*
2. 인터페이스는 만족스러운 유형에 대한 지식이 없어야한다.
인터페이스는 다른 유형이 해당 인터페이스의 멤버로 분류되는 데 필요한 사항을 정의해야한다.
예를 들어 자동차를 정의하는데 필요한 구성요소를 설명하는 인터페이스를 구축한다고 가정할 때
*/
type car interface {
	Color() string
	Speed() int
	IsFiretruck() bool
	//이렇게 하면 모든 차량의 종류를 car interface에 적어야한다.
	//IsSedan() bool , IsTank() bool 등을 계속 넣을 수는 없으니
	//하위 인터페이스를 만들면 좋다 파이어트럭에 대한 인터페이스를 만들면

}
type firetruck interface {
	car
	HoseLength() int
	//car 인터페이스의 Color와 Speed를 가져오면서도
	//Firetruck에 필요한 요소를 정의할 수 있다.
	//근데 이게 보아하니 약간 클래스 같은 느낌이...
}

/*
	3. 인터페이스는 클래스가 아니다.

	- 인터페이스는 클래스가 아니며 더 얇다(?).

	- 인터페이스에는 데이터를 생성하거나 삭제해야 하는 생성자 또는 해체자가 없다.

	- 인터페이스는 본질적으로 계층적이지 않지만 다른 인터페이스의 상위 집합이 되는
	  인터페이스를 만드는데 필요한 구문이 있다. (위의 car, firetruck 예시)

	- 인터페이스는 함수 서명을 정의하지만 기본 동작은 정의하지 않는다.
	  인터페이스를 만들면 구조체 메서드와 관련된 코드가 없는 경우가 많다...
	  위의 예시를 보면 Color() Speed() 전부 기본 동작이 정의되어있지 않다.

	  그래서 헷갈린다...^^... 나만 그런가...

	좋은 예시 모음들
	https://blog.boot.dev/golang/golang-interfaces/
*/
/* -------------------------------------------------------------- */
/*
	Q1. Interfaces should have as _ methods as possible
	인터페이스에는 가능한 한 _개의 메소드가 있어야 합니다.

	1. Complex
	2. Few
	3. Many

	A. 2
	약간의 메소드만 있는게 좋다 1번 정리 참고

*/

/*
	Q2. It's okay for types to be aware of the interfaces they satisfy
	유형이 자신이 만족하는 인터페이스를 인식하는 것은 괜찮습니다.

	1. True
	2. False

	A. 1
	2번 정리 참고..
*/

/*
	Q3. It's okay for interfaces to be aware of the types that satisfy them
	인터페이스가 자신을 만족시키는 유형을 인식하는 것은 괜찮습니다.

	1. False
	2. True

	A. 1
	2번 정리 참고
*/

/*
	Q4. Interfaces allow you to define a method's behavior once and use it for many different types
	인터페이스를 사용하면 메소드의 동작을 한 번 정의하고 이를 다양한 유형에 사용할 수 있습니다.

	1. False
	2. True

	A. 2
	이러려고 인터페이스가 있는거다
*/
