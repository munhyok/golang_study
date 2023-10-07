//Interfaces are Implemented Implicitly
/*
	인터페이스는 암시적으로 구현된다.

	type은 해당 메소드를 구현해 interface를 구현한다.
	다른 언어와는 달리 명시적인 표시가 없다 "implements" 키워드가 없다.

	암시적 인터페이스는 인터페이스의 정의와 그 구현을 분리한다.
	type에 메소드를 추가할 수도 있고 그 과정에서 자신도 모르게 다양한 인터페이스를
	구현할 수도 있지만 괜찮다.


	문제
	How is an interface fulfilled?
	인터페이스는 어떻게 구현됩니까?

	1. A type has all required interface's methods defined on it
	2. A struct embeds the interface in its definition

	Answer: 1


*/