package pointerTest

import "fmt"

func PointerTest() {
	a1 := 2
	b1 := a1
	fmt.Println(a1, b1) // 2, 2 출력됨
	a1 = 5
	fmt.Println(a1, b1) // 5, 2 출력됨 (b는 처음에 a의 값이였던 2를 담고 있음)

	// '&' 는 메모리를 참조 함
	a2 := 2
	b2 := &a2
	fmt.Println(a2, b2)  // 2, 0xc000a2000(변수 a의 메모리 주소) 출력됨
	fmt.Println(&a2, b2) // 0xc000a2000(변수 a의 메모리 주소), 0xc000a2000(변수 a의 메모리 주소) 출력됨

	// '*' 는 해당 메모리의 값을 볼 수 있음
	a3 := 2
	b3 := &a3
	fmt.Println(a3, *b3) // 2, 2 출력됨

	a3 = 5
	fmt.Println(a3, *b3) // 5, 5 출력됨 (변수 a3의 메모리 주소를 가르키던 b3 변수도 같이 값이 변경 됨)

	*b3 = 50
	fmt.Println(a3, *b3) // 50, 50 출력됨 (변수 a3의 메모리 주소를 가르키던 b3에서 값을 변경 가능)
}
