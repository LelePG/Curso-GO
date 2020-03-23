package main

import "fmt"

func main() {
	c:=f1(5)
	c()

}


func f1 (i int) func() {
	return func (){
		i = i*2
		fmt.Println(i)
}
}

