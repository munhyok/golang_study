//go Sturct Method
/*
	Go가 객체 지향 언어는 아니지만 구조체에 정의할 수 있는 메서드를 지원한다.
	메서드는 수신자가 있는 함수일 뿐이며, 수신자는 구문적으로 함수 이름 앞에 오는 특수 매개변수이다.

	수신자는 특별한 종류의 함수 매개변수일 뿐이다.
	수신자를 사용하면 구조체(및 기타유형)가 구현할 수 있는 인터페이스를 정의할 수 있기 때문에
	중요하다.

*/package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

// ?

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
