package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(b)
	r := 2 + rVal.Int()
	fmt.Println("rVal=", r)

	iV := rVal.Interface()
	// 将interface通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func main() {
	var num int = 100
	reflectTest01(num)
}
