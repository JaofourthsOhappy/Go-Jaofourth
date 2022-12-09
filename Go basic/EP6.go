package main

import "fmt"
//switch...case
func main(){
	var choice int
	fmt.Print("Enter choice:")
	fmt.Scanf("%d",&choice)

	switch choice{
	case 1:
		fmt.Println("Create Account")
	case 2:
		fmt.Println("deposit withdraw")
	default:
		fmt.Println("Not Successfully")
	}
}
