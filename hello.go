package main

import (
	"fmt"

	"github.com/goLangStudy/hiturl"
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

	/*
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

		account.Deposit(1000)
		fmt.Println(account.Balance())
		fmt.Println(&account)

		amount, err := account.Withdraw(2000)

		// error 처리
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed amount -", amount)
		}

		result, err := account.Withdraw(500)

		if err != nil {
			// log로 error 출력 후 프로그램 종료함.
			fmt.Println("failed amount -", amount)
			log.Fatalln(err)
		}

		fmt.Println(result)

		fmt.Println("====================")
		fmt.Println(account)
		fmt.Println("====================")
		fmt.Println(account1)
		fmt.Println("====================")
	*/

	/*
		// Dictionary

		// 이런식으로 타입만 선언해서 사용 가능 하지만 method를 활용하는게 더 좋음
		dictionary1 := mydict.Dictionary{"first": "first word"}

		fmt.Println(dictionary1)

		fmt.Println("========== Search Test ==========")

		definition1, err1 := dictionary1.Search("second")

		if err1 != nil {
			fmt.Println(err1)
		} else {
			fmt.Println(definition1)
		}

		definition2, err2 := dictionary1.Search("first")

		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(definition2)
		}

		fmt.Println("========== Add Test ==========")

		err3 := dictionary1.Add("first", "first word")

		if err3 != nil {
			fmt.Println(err3)
		} else {
			fmt.Println(dictionary1["first"])
		}

		err4 := dictionary1.Add("second", "second word")

		if err4 != nil {
			fmt.Println(err4)
		} else {
			fmt.Println(dictionary1["second"])
		}

		fmt.Println("========== Update Test ==========")

		err5 := dictionary1.Update("third", "third word")

		if err5 != nil {
			fmt.Println(err5)
		} else {
			fmt.Println(dictionary1["third"])
		}

		err6 := dictionary1.Update("second", "second word!!!")

		if err6 != nil {
			fmt.Println(err6)
		} else {
			fmt.Println(dictionary1["second"])
		}

		fmt.Println("========== Delete Test ==========")

		err7 := dictionary1.Delete("third")

		if err7 != nil {
			fmt.Println(err7)
		} else {
			fmt.Println(dictionary1["third"])
		}

		err8 := dictionary1.Delete("second")

		if err8 != nil {
			fmt.Println(err8)
		} else {
			fmt.Println(dictionary1["second"])
		}
	*/

	/*
		// hiturl
		var urls []string = []string{
			"https://www.naver.com",
			"https://www.daum.net",
			"https://www.google.com",
		}

		for _, url := range urls {
			err := hiturl.HitURL(url)

			if err != nil {
				fmt.Println(err)
			}

		}
	*/

	/*
		// channeltest
		arr := [5]string{"A", "B", "C", "D", "E"}

		for _, str := range arr {
			result := channeltest.Run(str)
			fmt.Println(result)
		}

		//channel 생성 - 변수명 := make(chan 체널이용해서 받을 타입)
		c := make(chan string)

		for _, str := range arr {
			go channeltest.RunChan(str, c)
		}

		// 채널에 arr의 length만큰 값을 넣었으므로, 그만큼 불러옴
		for i := range arr {
			// 변수에 채널에 담긴 값을 넣어줌
			// 프로그램 진행 중 '<-'를 만나면 값이 담기기 기다림 block
			result := <-c
			fmt.Println("Main response Channel : ", result, " - ", (i + 1))
		}
	*/

	/*
	 */
	// HitURL2 - Use Goroutine
	var urlArr [9]string = [9]string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.naver.com/",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// 채널 생성
	c := make(chan hiturl.ResponseResult)

	// urlArr 반복문
	for _, url := range urlArr {
		// goroutine을 이용한 병렬 호출
		go hiturl.HitURL2(url, c)
	}

	// 채널 값을 출력하기 위한 반복문
	for i := 0; i < len(urlArr); i++ {
		// goroutine 실행 한 만큼 채널의 값 가져오기
		fmt.Println(<-c)
	}
}
