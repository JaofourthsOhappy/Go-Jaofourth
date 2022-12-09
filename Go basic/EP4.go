package main

import "fmt"

//scanf
func main() {
	var name string
	var score int
	fmt.Print("Pls Enter name = ")
	fmt.Scanf("%s", &name)
	fmt.Print("Pls Enter score = ")
	fmt.Scanf("%d", &score)
	fmt.Println("hello",name)
	fmt.Println("Score + (พิเศษ)",score+10)
}