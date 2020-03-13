package main

import (
	"fmt"
)

func multiply (a, b int) (c int) {
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
}
