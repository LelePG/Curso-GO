package main

import "fmt"

type person struct {
	fname string
	lname string
	icecream []string
}
func main() {
	 p1:= person {
		fname: "lele",
		lname: "garcez",
		icecream: []string{"chocolate","brigadeiro"},
	}

	p2:= person {
		fname: "lele",
		lname: "pegoraro",
		icecream: []string{"nata","creme"},
	}

	fmt.Printf("%s %s gosta de ",p1.fname,p1.lname)
	for _, sorvete :=range p1.icecream{
		fmt.Printf("%s, ",sorvete)
	}
	fmt.Println()


	fmt.Printf("%s %s gosta de ",p2.fname,p2.lname)
	for _, sorvete :=range p2.icecream{
		fmt.Printf("%s, ",sorvete)
	}
	fmt.Println()
}
