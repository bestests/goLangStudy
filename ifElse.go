package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

func canIDrink2(age int) bool {
	if age < 18 {
		return false
	}

	return true
}

func canIDrink3(age int) bool {
	// if문 안에 변수 선언 가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink3(16))
}
