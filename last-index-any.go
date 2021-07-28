package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	LastIndexAny returns the index of the last instance of any Unicode code
	point from chars in s, or -1 if no Unicode code point from chars is present in s.
	*/
	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
}
