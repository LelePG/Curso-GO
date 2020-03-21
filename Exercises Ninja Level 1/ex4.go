package main

import "fmt"

type unicornio int

var x unicornio

func main() {
	fmt.Println(x)
	fmt.Printf("%T",x)
	x = 42
	fmt.Println(x)
}
