package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	Index returns the index of the first instance of substr in s,
	or -1 if substr is not present in s.
	*/
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}
