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
	const nameConst string = "Song"
	fmt.Println(nameConst)

	var nameVar string = "Song"

	nameVar = "Chang"

	fmt.Println(nameVar)

	// Only use in function
	name := "Sup"
	fmt.Println(name)

	resultNum1 := multiply(2, 2)

	fmt.Println(resultNum1)

	var resultNum2 int = multiply(2, 3)

	fmt.Println(resultNum2)

	var resultNum3 int = importTest.Plus(2, 3)

	fmt.Println(resultNum3)
}
