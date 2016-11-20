package main

import (
	"fmt"
	"time"
)

/** 
 * User: daixinyu
 * Date: 2016/11/20 
 * Time: 9:37 
 */

func show(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, "|", i)
	}
}

func no_channel() {
	for id := 0; id < 10; id++ {
		go show(id)
	}
	time.Sleep(time.Microsecond * 3)
}

func show_default(id int, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, "|", i)
	}
	ch <- 0
}

func default_channel() {
	ch := make(chan int)
	go show_default(1, ch)
	<-ch
}
func channel_mutil() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go fun1(ch1)
	go fun2(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}

func fun1(ch1 chan string) {
	for {
		ch1 <- "fun1"
		time.Sleep(3000)
	}
}

func fun2(ch1 chan string) {
	for {
		ch1 <- "fun2"
		time.Sleep(3000)
	}
}

func main() {
	//no_channel()
	//default_channel()
	channel_mutil()
}
