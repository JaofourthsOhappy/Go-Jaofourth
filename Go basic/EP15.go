package main

import "fmt"
type Product struct{
	name string
	price float64
	category string
	discount int
}
func main(){
	//stuct
	product1 := Product{name: "Pencil",price: 15.50,category: "study",discount: 5}
	fmt.Println(product1)
	fmt.Println(product1.name)
	product1.price = 20
	fmt.Println(product1.price)
	fmt.Println(product1.category)
	fmt.Println(product1.discount)
  
}  