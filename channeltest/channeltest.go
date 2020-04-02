package channeltest

import (
	"fmt"
	"time"
)

// Run - not use goroutine
func Run(str string) bool {
	time.Sleep(time.Second * 2)
	fmt.Println("Run - str : ", str)
	return true
}

// RunChan - use goroutine
/*
	채널을 변수로 받을땐, 이 채널을 이용해 전달할 타입도 같이 써 줌
*/
func RunChan(str string, c chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("RunChan - str : ", str)

	// 채널에 값을 넣어 줌
	c <- ("RunChan channel value : " + str)
}
