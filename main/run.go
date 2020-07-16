package main
import (
	"fmt"
)

func main()  {
	angel := "Heros never die"
	angleBytes := []byte(angel)
	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))
}