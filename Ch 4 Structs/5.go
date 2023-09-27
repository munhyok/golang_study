// go embedded structs
/*
	Go는 객체 지향 언어가 아니다.
	하지만 포함된 구조체는 때때로 유용할 수 있는 일종의 데이터 전용 상속을 제공한다.
	Go는 완전한 의미에서 클래스나 상속을 지원하지 않지만 Embedded Structs는
	구조체 정의 간에 필드를 높이고 공유하는 방법 중 하나이다.


	Embedded vs Nested
	포함된 구조체의 필드는 중첩된 구조체와 달리 최상위 수준에서 액세스된다.
	승격된 필드는 복합 리터럴에서 사용할 수 없는 것을 제외하고 다른 필드 처럼 액세스할 수 있다.

*/

package main

import "fmt"

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	//Embedded Structs 인스턴스화
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
