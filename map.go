package main

import (
	"fmt"
	"strings"
)
//decagonedo: Coding@edo123
func main()  {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	//abc := "wertyu"
	fmt.Println(rot13(700000))
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
//length
//123	321