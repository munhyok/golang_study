//early return
//함수에서 일찍 반환하는 기능
//중접된 조건부를 간단하게 작성할 수 있다

//returnGuard Clauses는 함수에서(또는 continue루프를 통해) 조기에 중첩된 조건문을 1차원으로 만드는 기능을 활용
//if/else 체인을 사용하는 대신 각 조건부 블록 끝에 있는 함수에서 조기에 반환한다

//중첩 조건부 예시

package main

func getInsuranceAmount_ifelse(status insuranceStatus) int {
	amount := 0

	if !status.hasInsurance() {
		amount = 1
	} else {
		if status.isTotaled() {
			amount = 10000
		} else {
			if status.isDented() {
				amount = 160
				if status.isBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}

//early_return으로 정리

func getInsuranceAmount_earlyReturn(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}

//간결해서 장점만 있을까
//Early Return의 단점
//함수의 반환이 여러곳으로 흩어지게 되어 함수의 복잡도를 높이고
//오히려 가독성이 떨어질 수 있다
//간결한 코드에선 오히려 if-else문이 더 가독성이 좋을 것 같다
