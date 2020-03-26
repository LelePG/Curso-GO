package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Programa principal")
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()

}

func foo(){
	fmt.Println("FOO")
	wg.Done()
}

func bar(){
	fmt.Println("BAR")
	wg.Done()
}
