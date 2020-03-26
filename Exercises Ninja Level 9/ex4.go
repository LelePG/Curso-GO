package main

import (
	"fmt"
	"runtime"
	"sync"
)

//rodar com go run -race ex3.go pra exibir a data race
var acumulador int //já inicia com zero
var wg sync.WaitGroup
var m sync.Mutex

func main() {
	valor :=5
	wg.Add(valor)
	for i:=0;i<valor;i++{
		go func (){
			m.Lock()
			ac:=acumulador
			//runtime.Gosched()
			ac++
			acumulador = ac
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			m.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("acumulador: ",acumulador)
}


