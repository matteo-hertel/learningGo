package main

import "fmt"

func main(){
	fmt.Print("Convert Fahrenheit To Celsius: ")
	var input float32
	fmt.Scanf("%f", &input)
	
	output := fahrenheitToCelsius(input)
	fmt.Println(output)
}

func fahrenheitToCelsius(fahrenheit float32) float32{
	return (fahrenheit - 32) * 5/9
}
