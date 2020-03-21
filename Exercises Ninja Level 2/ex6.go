package main

import "fmt"

const (
	ano1 = 2020 +iota
	ano2 = 2020 + iota
	ano3 = 2020 + iota
	ano4 = 2020 + iota
)
func main() {
	fmt.Println(ano1,ano2,ano3,ano4)
}
