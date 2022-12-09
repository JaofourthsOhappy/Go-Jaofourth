package main

import "fmt"

func main(){
	  var numbers1 [3] int = [3]int{100,200,300}
	  numbers2 := [3]int{100,200,300}
	  numbers3 := [...]int{1,2,3,4,5,6,7,8,9,10,11,12}
	  fmt.Println(numbers1)
	  fmt.Println(numbers2[1])
	  fmt.Println(numbers3,len(numbers3))

}