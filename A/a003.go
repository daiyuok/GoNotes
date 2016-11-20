package main

import "fmt"

/**
 * User: daixinyu
 * Date: 2016/11/19
 * Time: 22:25
 */



type user struct {
	id   int
	name string
	age  int
}

type family struct {
	user
	num int
}

func (u user) call() string {
	fmt.Printf("%+v \n", u)
	return fmt.Sprintf("%+v", u)
}

func (u *user)area()int  {
	u.id=2
	return u.id*u.age
}

func main() {
	baba := user{1, "爸爸", 50}
	f1 := family{baba, 3}
	fmt.Println(f1.call())
	fmt.Println(f1.area())
	fmt.Println(f1.call())
}
