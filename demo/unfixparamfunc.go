package main

import "fmt"

func sum1(nums ...int){
	fmt.Printf("%T\n",nums)
	total := 0
	for _,num := range nums{
		total += num
	}

	fmt.Println(total)
}

func main(){
	sum1(1,2,3)
	sum1(1,2,3,4,5)

	var nums = []int{1,2,3,4}
	sum1(nums...)
}