/*
	Name Your Interface Arguments

	인터페이스 인수 이름 지정하기

	아래 예시 처럼 명확하게 하기 위해 인수 이름 추가하고 데이터 반환이 가능하다
	그냥 type만 적어두면 코드만 읽었을 때 뭔소리인지 모르겠으니..
	저렇게 이름을 추가해서 보기 좋게 만들자 :)

	Q. Are you required to name the arguments of an interface in order for your code to compile properly?
	코드를 올바르게 컴파일하려면 인터페이스 인수의 이름을 지정해야 합니까?

	1. Yes
	2. No

	A. 2
	굳이 지정하지 않아도 된다.
	단지 보기 좋게 만들어야 나중에 협업하기 좋아지니...

*/

package main

type Copier interface {
	Copy(string, string) int
}

type Copier_Name interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
