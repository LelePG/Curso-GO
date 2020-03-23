package main

import "fmt"

func main() {
	resultado:=fa(fb(4),2)
	fmt.Println(resultado)
}


func fa (fb func (int) int , i int) int {//fa incrementa o valor passado
	return fb(i)+1//incrementa o duplicado
}

func fb(i int) int  {//fb duplica
	return i*2

}