package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	IndexAny returns the index of the first instance of any Unicode code
	point from chars in s, or -1 if no Unicode code point from chars is
	present in s.
	 */
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
}
