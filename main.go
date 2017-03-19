package main

import "fmt"

const tempcTof float64 = 1.8

func main() {
	var tempC float64
	fmt.Print("Enter temp in C: ")
	fmt.Scan(&tempC)
	tempF := tempC*tempcTof + 32
	fmt.Println(tempC, " degrees C is ", tempF, " degrees F ")
}
