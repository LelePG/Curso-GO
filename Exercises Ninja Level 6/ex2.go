package main

import "fmt"

func main() {
	x:= []int {1,2,3,4,5,6}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(v ...int) int {
	total:=0
	for _,c:=range v{
		total +=c
	}
	return total
}

func bar(v []int) int{
	total:=0
	for _,c:=range v{
		total +=c
	}
	return total
}