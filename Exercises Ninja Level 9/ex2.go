package main

import "fmt"

type person struct{
	First string
	Age int
}

func (p *person) speak(){
	fmt.Println("I'm ",(*p).First)
}

type human interface {
	 speak()
}

func saySomething(h human){
	h.speak()
}


func main() {
lele:=person{
	First: "Lele",
	Age:   19,
}

saySomething(&lele)//É passado um ponteiro pra um humano, que é uma pessoa
//saySomething(lele)//speak method has a pointer reciever. Isso não funciona.

}