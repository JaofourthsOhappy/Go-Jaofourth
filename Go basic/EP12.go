package main

import "fmt"

func main(){
	showMessage()
	showMessageAndName("fourth")
	showMessageAndName("amm")
	total(30,21)
	delivery := getDelivery()
	fmt.Println("Delivery fee = ",delivery)
	myCart := getTotalCart(500,700)
	fmt.Println("MyCart = ", myCart)
}

func showMessage(){
	fmt.Println("Hello Jaofourth")
}

func showMessageAndName(name string){
	fmt.Println("Hello",name)
}

func total(number1 int,number2 int){
	fmt.Println("Total = ",number1+number2)
}

func getDelivery() int{
	return 50
}

func getTotalCart(number1,number2 int)int{
	total := number1 + number2
	return total
}
