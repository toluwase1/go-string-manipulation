package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, false)
	fmt.Println(string(b))
}