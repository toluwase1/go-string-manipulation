package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s
	 */
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))
}
