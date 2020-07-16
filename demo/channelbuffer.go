package main

import "fmt"

func main(){
	message := make(chan string,2)

	message <- "hello"
	message <- "world"
	message <- "goalng"

	fmt.Println(<-message)
	fmt.Println(<-message)
	// fmt.Println(<-message)
	// fmt.Println(<-message)
}