package main

import "fmt"

// ! chan<- 이렇게 쓰면, 채널로 뭔가를 보내기만 가능하다.
func sendOnly(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf("send %d", i)
		c <- i
	}
}

// ! <-chan 이렇게 쓰면 채널로 뭔가를 받기만 가능하다.
func receiveOnly(c <-chan int) {
	for {
		data, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("receive %d", data)
	}
}

func goRoutines() {
	// ! goroutine을 사용하고 통신하려면 채널이 필요하다. 따라서, 채널을 생성한다.
	c := make(chan int)
	// ! goroutine을 사용하는 방법
	go sendOnly(c)
	receiveOnly(c)
}
