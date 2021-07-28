package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	EqualFold reports whether s and t, interpreted as UTF-8 strings,
	are equal under Unicode case-folding, which is a more general
	form of case-insensitivity.
	 */
	fmt.Println("equal fold")
	fmt.Println(strings.EqualFold("Go","go0"))
	//var a float64 = -1
	//var b float64 = 2
	//
	//fmt.Println(a/b)
}




