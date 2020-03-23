package main

import "fmt"

func main() {
	fmt.Println("Inicio")
	defer foo1()
	fmt.Println("Meio")
	defer foo2()
	fmt.Println("Fim")

}

func foo1(){
	fmt.Println("Foo1")
}

func foo2(){
	fmt.Println("foo 2")
}
