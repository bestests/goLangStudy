package arraySlice

import (
	"fmt"
)

func ArraySlice () {

	/* Array */

	// 배열은 크기가 정해져 있음
	// 배열 선언 - 변수명 := [크기]타입{값...}
	strArr := [3]string{"a", "b", "c"}

	// 배열 사용 - 변수명[인덱스]
	fmt.Println(strArr[0])

	// 배열 변경 - 변수명[인덱스] = 값
	strArr[0] = "d"
	fmt.Println(strArr[0])

	// 크기가 정해져있으므로, 크기 이상의 인덱스 사용 시 error
	// strArr[3] = "e"

	/* Slice */

	// 크기가 정해져있지않음
	// slice 선언 - 변수명[]타입{값...}
	strSlice := []string{"A", "B", "C"}

	// slice 사용
	fmt.Println(strSlice[0])

	// slice 변경
	strSlice[0] = "D"
	fmt.Println(strSlice[0])

	// 크기, 용량 출력
	fmt.Println(len(strSlice), cap(strSlice))

	// slice 값 추가
	// append 함수 사용
	// append(slice, value)
	// 기존 slice에 추가가 아닌 return을 함
	strSlice = append(strSlice, "E")

	// 변경 후 값 확인
	fmt.Println(strSlice)
	fmt.Println(len(strSlice), cap(strSlice))
}