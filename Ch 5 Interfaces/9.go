/*
	Name Your Interface Arguments

	인터페이스 인수 이름 지정하기

	아래 예시 처럼 명확하게 하기 위해 인수 이름 추가하고 데이터 반환이 가능하다
	그냥 type만 적어두면 코드만 읽었을 때 뭔소리인지 모르겠으니..
	저렇게 이름을 추가해서 보기 좋게 만들자 :)

	Q. Why would you name your interface's method's parameters?
	인터페이스의 메소드 매개변수 이름을 지정하는 이유는 무엇입니까?

	1. Execution Speed
	2. Memory Savings
	3. Readability & clarity

	A. 3

	그냥 읽기 좋아야 나중에 수정하기 편하잖아...

*/

package main

type Copier interface {
	Copy(string, string) int
}

type Copier_Name interface {
	Copy(sourceFile string, destinationFile string) (bytesCopied int)
}
