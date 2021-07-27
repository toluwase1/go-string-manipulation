package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("contains Any")
	fmt.Println(strings.ContainsAny("thomas", ""))
	fmt.Println(strings.ContainsAny("thomas", "sg"))
	fmt.Println(strings.ContainsAny("thomas", "0"))
	fmt.Println(strings.ContainsAny("thomas","i"))
	fmt.Println(strings.ContainsAny("",""))
}