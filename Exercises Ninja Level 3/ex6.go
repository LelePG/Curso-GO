package main

import "fmt"

func main() {
	for c:=10; c<=100; c++{
		if c%5 == 0{
			fmt.Println(c)
		}
	}
}