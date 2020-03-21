package main

import "fmt"

func main() {
	 a := [5]int {1,2,3,4,5}//declaração usando string literal

	for _, v:= range a{//a variável v toma os valores dentro de a. O valor descartado, é o índice
		fmt.Println(v)
	}

	fmt.Printf("Tipo: %T",a)
	fmt.Println()
}
