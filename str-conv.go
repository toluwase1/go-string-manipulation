package main

import (
	"fmt"
	"strconv"
)

func main() {
	price:= "The price of rice is " + strconv.Itoa(150) + " naira"
	//more accurate
	price1:= "The price of beans is " + strconv.FormatFloat(20, 'f', 2, 64) + " naira"

	divisor:=2
	var1:= 20.02
	fmt.Printf("%T", var1)
	fmt.Println()
	fmt.Println(price)
	fmt.Println(price1)

	fmt.Println(var1/float64(divisor))


}
