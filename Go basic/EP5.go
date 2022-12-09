package main

import "fmt"
//if...else
func main(){
	var score int
	fmt.Print("Enter score:")
	fmt.Scanf("%d",&score)
	if score >= 50 {
		println("You pass")
	}else{
		println("You not pass")
	}
}