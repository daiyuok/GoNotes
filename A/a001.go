package main

import (
	"fmt"
)

func main() {
	var a = 1
	if a < 0 {
		fmt.Println("0")
	} else if a > 0 {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
}
