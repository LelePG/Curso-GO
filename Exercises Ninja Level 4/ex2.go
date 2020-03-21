package main

import "fmt"

func main() {
	slc := []int {12,23,34,45,56}
	for _, value := range slc{
		fmt.Println(value)
	}

	fmt.Printf("Tipo: %T", slc)
	fmt.Println()
}
