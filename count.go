package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("count")
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
	fmt.Println(strings.Count("eeeeee", "ee"))
}
