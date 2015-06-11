package main

import "fmt"

func main() {
	var tempInFahrenheit float32
	fmt.Scanf("%d", tempInFahrenheit)
	tempInCelsius := 5 * (tempInFahrenheit - 32) / 9
	var result string = "" + tempInFahrenheit + " Fahrenheit is " + tempInCelsius + " Celsius"
	fmt.Print(result)
}
