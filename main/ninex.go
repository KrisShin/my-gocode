package main

import "fmt"

func ninex()  {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d ", i,j,i*j)
		}
		fmt.Println()
	}
}