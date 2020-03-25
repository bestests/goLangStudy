package main

import (
	"fmt"

	"github.com/goLangStudy/accounts"
)

func multiply(a, b int) (c int) {
	c = a * b
	return
}

func main() {
	/*
		// import test
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
		// import test
	*/

	/*
		// 2. forRangeArgs
		var forRangeArgsResult = forRangeArgs.SuperAdd(1, 2, 3, 4)

		fmt.Println(forRangeArgsResult)
		// 2. forRangeArgs
	*/

	/*
		// 3. ifElse
		fmt.Println(ifElse.CanIDrink(16))
		fmt.Println(ifElse.CanIDrink2(16))
		fmt.Println(ifElse.CanIDrink3(16))
		// 3. ifElse
	*/

	/*
		// 4. switch
		fmt.Println(switchTest.SwitchTest1(16))
		fmt.Println(switchTest.SwitchTest2(16))
		fmt.Println(switchTest.SwitchTest3(16))
		fmt.Println(switchTest.SwitchTest4(16))
		// 4. switch
	*/

	/*
		// 5. pointer
		pointerTest.PointerTest()
	*/

	/*
		// 6. arraySlice
		arraySlice.ArraySlice()
	*/

	/*
		// 7. map
		mapTest.MapTest()
	*/

	/*
		// 8. structs
		structs.Structs()
	*/

	/**/
	// account
	account := accounts.NewAccount("kim")
	fmt.Println(&account)
	fmt.Println(*account)

	// owner는 private이므로 접근 못함.
	//fmt.Println(account.owner)

	account = accounts.NewAccount("lee")
	fmt.Println(&account)
	fmt.Println(*account)

	account1 := accounts.NewAccount("park")
	fmt.Println(&account1)
	fmt.Println(*account1)
}
