package main

import "fmt"

func main() {
	c:=make(chan int, 100)
	go preenche(c)
	mostra(c)
}


func preenche(ch chan int){
	for i:=0;i<100;i++{
		ch<-i
	}
	close(ch)
}

func mostra(ch chan int){
	for c := range ch{
		fmt.Println(c)
	}
}