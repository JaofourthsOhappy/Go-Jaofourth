package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(40,30))
}

func add(x int,y int)int{
	return x + y
}