package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func main() {
	p:= person{
		first:"Lele",
		last:"PG",
		age:19,
	}
	p.speak()
}

func (p person) speak (){
	fmt.Printf("Meu nome é %s %s tenho %d anos.", p.first,p.last,p.age)
	fmt.Println()
}