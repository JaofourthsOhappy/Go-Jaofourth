package main

import "fmt"

func main(){
	//slice
	numbers := []int{100,200,300}
	fmt.Println(numbers)
	numbers = append(numbers,400)
	fmt.Println(numbers)
	fmt.Println(numbers[:1])

}