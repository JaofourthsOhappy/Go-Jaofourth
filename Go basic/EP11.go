package main

import "fmt"
func main(){
	numbers := []int{10,20,30,40,50,60,70,80,90}
	language := map[string]string{"TH":"Thailand","JP":"Japan"}
	for index := 0;index<len(numbers);index++{
		fmt.Println(numbers[index])
	}
	for index , value := range numbers{
		fmt.Println(index,value)
	}
	for _, value := range numbers{  
		fmt.Println(value)
	}
	for _, value := range language{  
		fmt.Println(value)
	}
}