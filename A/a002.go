package main

/**
 * User: daixinyu
 * Date: 2016/11/19
 * Time: 22:25
 * 切片----动态数组
 */

import "fmt"

func main() {
	x := make([]int, 0, 5)
	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
	for i := 0; i < 8; i++ {
		println(x[i])
	}
}
