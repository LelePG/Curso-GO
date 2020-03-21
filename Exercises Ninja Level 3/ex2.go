package main

import "fmt"

func main() {
	for num := 65; num < 91; num++{
		fmt.Println(num)
		for c:=0;c<3;c++{
			fmt.Printf("%#U",num)
			fmt.Println()
		}
	}
}
