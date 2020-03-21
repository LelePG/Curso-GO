package main

import "fmt"

type person2 struct {
	fname string
	lname string
	icecream []string
}
func main() {
	p1:= person2 {
		fname: "lele",
		lname: "garcez",
		icecream: []string{"chocolate","brigadeiro"},
	}

	p2:= person2 {
		fname: "lele",
		lname: "pegoraro",
		icecream: []string{"nata","creme"},
	}

	m := map[string] person2{
		p1.lname: p1,
		p2.lname: p2,
	}


	for _, pessoa := range m{
		fmt.Println(pessoa.fname, pessoa.lname)
		for _, sorvete :=range pessoa.icecream{
			fmt.Printf("%s, ",sorvete)
		}
		fmt.Println()

	}
	fmt.Println()

}
