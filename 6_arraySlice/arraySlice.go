package arraySlice

import (
	"fmt"
)

func ArraySlice () {

	/* Array */

	// 배열 선언 - 변수명 := 타입[크기]{"값1", ...}
	strArr := [3]string{"a", "b", "c"}

	// 배열 사용
	fmt.Println(strArr[0])

	// 배열 변경
	strArr[0] = "d"
	fmt.Println(strArr[0])

	// 크기가 3이므로 error
	// strArr[3] = "e"

	/* Slice */
	// 배열과 달리 크기에 제한없음
	// 크기 초과 시 자동으로 늘어남
	
	strSlice := []string{"A", "B", "C"}

	// slice 사용
	fmt.Println(strSlice[0])
}