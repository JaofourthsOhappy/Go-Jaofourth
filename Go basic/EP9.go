package main

import "fmt"

func main(){
	//map
	country := map[string]string{"TH":"Thailand","JP":"Japan"}
	fmt.Println(country["TH"])

	value, check := country["JP"]

	if check {
		fmt.Println(value)
	}else{
		fmt.Println("No result")
	}

}

