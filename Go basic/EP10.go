package main

import "fmt"

func main(){
	// for count:=1 ; count<=3; count++{
	// 	fmt.Println("hello go programming")
	for count := 10 ; count >= 1; count--{
		fmt.Println(count)
		if count == 5{
			break
		}
	}
	fmt.Println("End")
}