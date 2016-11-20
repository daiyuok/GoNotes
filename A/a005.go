package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strconv"
)

/** 
 * User: daixinyu
 * Date: 2016/11/20 
 * Time: 9:20 
 */

func main() {
	steam, _ := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\aaa.txt")
	fmt.Println(string(steam))

	file, _ := os.Open("C:\\Users\\Administrator\\Desktop\\aaa.txt")
	buff := make([]byte, 1024)
	n, _ := file.Read(buff)
	fmt.Println(string(buff[:n]))

	fmt.Println("what is your name:\n")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)

	num0 := 1
	num1 := 1.5
	str0 := "100"
	str01 := "100.1"

	fmt.Println(float64(num0))
	fmt.Println(int(num1))
	fmt.Println(strconv.ParseInt(str0, 0, 64))
	fmt.Println(strconv.ParseFloat(str01, 64))

}
