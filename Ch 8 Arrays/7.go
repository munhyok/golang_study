/*
	Len and Cap Review

	6.go 에서 len cap 정리한 부분 참고


	해당 코드는 slice에 data를 추가하는 함수인데
	데이터가 용량을 초과하면 make함수를 이용해 slice의 cap를 재할당해
	cap을 늘려 slice를 출력하는 함수이다
	이렇게 cap의 용량보다 data의 수가 더 많을 때 make()를 이용해 cap을 더 늘릴 수 있다.

	func Append(slice, data []byte) []byte {
	    l := len(slice)
	    if l + len(data) > cap(slice) {  // reallocate
	        // Allocate double what's needed, for future growth.
	        newSlice := make([]byte, (l+len(data))*2)
	        // The copy function is predeclared and works for any slice type.
	        copy(newSlice, slice)
	        slice = newSlice
	    }
	    slice = slice[0:l+len(data)]
	    copy(slice[l:], data)
	    return slice
	}

*/

// 문제

/*
	Q1. What does the cap() function return?

	1. The last element of the slice
	2. The maximum length of the slice before reallocation of the array is necessary

	A. 2

	Q2. What does the len() function return?

	1. The current length of the slice
	2. The maximum length of the slice before reallocation of the array is necessary

	A. 1

	Q3. What do len() and cap() do when a slice is nil?

	1. Return 0
	2. Panic

	A. 1
*/
