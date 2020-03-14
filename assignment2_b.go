package main

import(
    "fmt"
)

//var number1, number2, number3 float
func performArtithmetic(){
	fmt.Print("Enter First Number :")
	var number1 int
	fmt.Scanln(&number1)	 //store the first input

	fmt.Print("Please Enter the Second Number Greater than 0 \n")
	fmt.Print("Enter Second Number :")
	var number2 int
	fmt.Scanln(&number2)	//store the scond input

	var number3 int
	// Addition
	number3 = number1+number2
	fmt.Println("Addition of Two Number is : ",number3)

	// Subtraction
	number3 = number1-number2
	fmt.Println("Subtraction of Two Number is : ",number3)

	// Mulltiplication
	number3 = number1*number2
	fmt.Println("Multiplication of Two Number is : ",number3)

	// Division
	number3 = number1/number2
	fmt.Println("Division of Two Number is : ",number3)
}

func main(){
	//first line print stmt
	fmt.Println("Simple Calculator Program using Defer Function")
	fmt.Println("-----------------------------------------------")

	//calling defer function
	defer fmt.Println("-----------------------------------------------")
	defer fmt.Println("End of the Calculator Programming!!")

	//getting input from user and ready to performed Calculator function
	fmt.Println("Please fill the Following Informations to Continue the Actions:")

	performArtithmetic()
}