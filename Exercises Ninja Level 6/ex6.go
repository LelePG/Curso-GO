package main

import "fmt"

func main() {
	x := 5

	a:= func(i int) int{//Assinatura da função
		v:= i+1
		return v
	}(x)//Passagem de parâmetros

fmt.Println(a)
}


