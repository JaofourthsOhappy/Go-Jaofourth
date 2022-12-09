package main 

import "fmt"

func main(){
	result := summation(10,20,30,40,50,60,70)
	fmt.Println("Result =",result)
}

func summation(numbers ... int)int{
	total := 0
	for _,value := range numbers{
		total += value
	}
	return total
}