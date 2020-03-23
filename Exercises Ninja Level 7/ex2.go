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

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	(*p).first = "James"
	(*p).last = "Bond"
	(*p).age = 34
}