package main

import "fmt"


func main(){
	result,check := sumnumber(100,400)
	fmt.Println("Sum = ",result)
	fmt.Println(check)
}

func sumnumber(number1 , number2 int)(int,string){
	total :=  number1 + number2
	status := ""
	if total%2 == 0{
		status = "even"
	}else{
		status = "odd"
	}
	return total , status
}