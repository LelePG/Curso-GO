package main

import "fmt"

type unicornio int

var x unicornio
var y int
func main() {
	fmt.Println(x)
	fmt.Printf("%T",x)
	x = 420
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T",y)

}
