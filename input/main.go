package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Println("Please enter temp in Fahrenheit:")
	fmt.Scan(&fahrenheit)
	fmt.Println(fahrenheit, "in Celcius is: ", convertToCelsius(fahrenheit))
}

func convertToCelsius(input float64) float64 {
	celsius := (input - 32) * 5 / 9
	return celsius
}
