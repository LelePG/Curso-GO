package main

import (
	"fmt"
)

type customErr struct{
	nome string
}

func (err customErr) Error() string{
	return fmt.Sprint(err.nome)
}

func foo(er customErr){
	fmt.Println(er)
}


func main() {
	c1:=customErr{nome:"Out of money"}
	foo(c1)
}
