package main

import "fmt"

func main(){
	var number1 int = 10
	var number2 int = 6
	fmt.Println("If + =",number1+number2)
	fmt.Println("If - =",number1-number2)
	fmt.Println("If * =",number1*number2)
	fmt.Println("If / =",number1/number2)

	number3,number4 := 20,12
	fmt.Println("Are =",number3==number4)
	fmt.Println("Are !=",number3==number4)
	fmt.Println("Are >",number3>number4)
	fmt.Println("Are <",number3<number4)
	fmt.Println("Are >=",number3>=number4)
	fmt.Println("Are <=",number3<=number4)
}