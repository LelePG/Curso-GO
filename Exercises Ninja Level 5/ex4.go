package main

import "fmt"

func main() {
	p1:= struct {
		fname string
		lname string
		icecream []string
	} {
		fname: "lele",
		lname: "garcez",
		icecream: []string{"chocolate","brigadeiro"},
	}


	fmt.Printf("%s %s gosta de ",p1.fname,p1.lname)
	for _, sorvete :=range p1.icecream{
		fmt.Printf("%s, ",sorvete)
	}
	fmt.Println()



}
