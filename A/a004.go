package main

import "fmt"

/**
* User: daixinyu
* Date: 2016/11/19
* Time: 22:41
生产者消费者模式
*/

func consume(data chan int, done chan bool) {
	for x := range data {
		fmt.Printf("rec:%d \n", x)
	}
	done <- true
}

func pub(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func main() {
	data := make(chan int)
	done := make(chan bool)
	go pub(data)
	go consume(data, done)
	<-done
}
