package main

import "fmt"

type rect struct {
	width,height int
}

func (r *rect) area() int {
	return r.width*r.height
}

func (r rect) perim() int {
	return (r.width+r.height)*2
}

func main(){
	r := rect{width:5, height:20}
	
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r

	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
