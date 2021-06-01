package main

import (
	"Golang/myfunc"
	"fmt"
)

func main() {
	//初始化
	myfunc.SayHello()
	//函数类型作为函数变量
	ret2 := myfunc.Calc(10, 20, myfunc.Add)
	fmt.Println(ret2)
	a, _ := myfunc.Do("+")
	fmt.Println(a(1, 2))
}
