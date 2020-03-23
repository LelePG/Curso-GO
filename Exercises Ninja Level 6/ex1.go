package main

import "fmt"

func main() {
	 x := 5

	 x = foo(x)
	 fmt.Println(x)

	  x,  s := bar(x ,"Lele")
	 fmt.Println(x ,s)

}

func foo (i int ) int{
	return i+1
}

func bar (i int , nome string ) (int, string){
	return i+1 , "Olá "+ nome
}
