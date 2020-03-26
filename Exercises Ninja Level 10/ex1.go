package main

import (
"fmt"
)

func main() {
	//c := make(chan int)//a
	c:= make (chan int, 1)//b
	//go func (){//a
		c <- 42
	//}()//a
	fmt.Println(<-c)
}
