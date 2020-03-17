package main

import (
	"fmt"

	"github.com/goLangStudy/importTest"
)

func multiply(a, b int) (c int) {
	c = a * b
	return
}

func main() {
	const nameConst string = "Test"
	fmt.Println(nameConst)

	var nameVar string = "Name"

	nameVar = "Name!!!"

	fmt.Println(nameVar)

	// Only use in function
	name := "name"
	fmt.Println(name)

	resultNum1 := multiply(2, 2)

	fmt.Println(resultNum1)

	var resultNum2 int = multiply(2, 3)

	fmt.Println(resultNum2)

	var resultNum3 int = importTest.Plus(2, 3)

	fmt.Println(resultNum3)

	var resultNum4 int = importTest.Plus(3, 4)

	fmt.Println("resultNum4 - ", resultNum4)

	var resultNum5 int = importTest.Plus(4, 7)

	fmt.Println("resultNum5 - ", resultNum5)
}
