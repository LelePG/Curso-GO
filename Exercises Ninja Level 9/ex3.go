package main

import (
	"fmt"
	"runtime"
	"sync"
)

//rodar com go run -race ex3.go pra exibir a data race
var acumulador int //já inicia com zero
var wg sync.WaitGroup
func main() {
	valor :=5
	wg.Add(valor)
	for i:=0;i<valor;i++{
		go func (){
			ac:=acumulador
			runtime.Gosched()
			ac++
			acumulador = ac
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(acumulador)
}


