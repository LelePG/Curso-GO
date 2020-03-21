package main

import "fmt"

func main() {
	num:= 0
	if num == 0{
		fmt.Println("Nao é impar nem par")
	} else if num%2==0{
		fmt.Println("Numero par")
	} else{
		fmt.Println("numero impar")
	}
}
