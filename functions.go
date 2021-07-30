package main

import "fmt"

func main() {

	//abc:= func printHello () {
	//	fmt.Println("hello")
	//}
	test:= func() {
		fmt.Println("hello")
	}
	test()


	abc:= func () {
		fmt.Println("hello")
	}
	abc()

	//cde:= func(x string) int{
	//	return x*1
	//} (x)
	//ninja


}

func abc (s string) string{
	return s
}

func SimpleFunction() {
	fmt.Println("Hello World")
}

