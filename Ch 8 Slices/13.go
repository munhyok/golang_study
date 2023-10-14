/*
	Tricky Slices

	append()는 매개변수의 기본 배열을 변경하고 새 슬라이스를 반환한다.

	// dont do this! 이런 행위
	someSlice = append(otherSlice, element)
	이런식으로 다른 매개변수를 이용한 append는 하지말자

	//좋은 예시
	names = append(names, "kang")


	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4)

*/

// 문제

/*
	Q1. Why is 5 the final value in the last index of 'j'?

	1. The Go team be trollin'
	2. j and g point to same underlying array so g's append overwrote j
	3. Because append only works properly when the number of elements is < 10

	A. 2


	Q2. Why doesn't the bug regarding slices 'j' and 'g' in example 2 occur in example 1 as well?

	1. Because there are fewer elements and Go's runtime can't handle more than ~8
	2. The array's cap() is exceeded so a new underlying array is allocated

	A. 2

	Q3. How can you best avoid these kinds of bugs?

	1. Don't use the append() function
	2. Always assign the result of the append() function back to same slice
	3. Always assign the result of the function to a new slice

	A. 2


*/