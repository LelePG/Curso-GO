package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//rodar com go run -race ex3.go pra exibir a data race
var acumulador int64 //já inicia com zero
var wg sync.WaitGroup
func main() {
	valor :=5
	wg.Add(valor)
	for i:=0;i<valor;i++{
		go func (){
			atomic.AddInt64(&acumulador,1)
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(acumulador)
}


