package main

import "fmt"

func main() {
	fmt.Println(fz())
	fmt.Println(fz())
	fmt.Println(fz())

}

func fz() int {
	z := 0
	z++
	return z
}
