package main

import (
	"fmt"
	"strings"
)

func main()  {

/* lexicographical comparisms
	Letter	ASCII Code	Binary
   	A		065			01000001
   	B		066			01000010
 */
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))


}
