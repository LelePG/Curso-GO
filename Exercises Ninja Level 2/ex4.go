package main

import "fmt"

func main() {
	v1 := 10
	fmt.Printf("Decimal:%d--Hexa:%x--Binario:%b",v1,v1,v1)
	v1 = v1<<1
	fmt.Println()
	fmt.Printf("Decimal:%d--Hexa:%x--Binario:%b",v1,v1,v1)
	fmt.Println()

}
