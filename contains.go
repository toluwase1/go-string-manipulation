package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("contains")
	fmt.Println(strings.Contains("foobar", "foo"))
	fmt.Println(strings.Contains("foobar","abo"))
	fmt.Println(strings.Contains("foo", "foo"))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("foobar "," "))
}
