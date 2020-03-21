package main

import "fmt"

func main() {
	slc :=[][]string{[]string{"James", "Bond", "Shaken, not stirred"},[]string{"Miss", "Moneypenny", "Helloooooo, James."} }

	for _,valor := range slc{
		for _, elemento := range valor{
			fmt.Printf("%s--",elemento)
		}
		fmt.Println()
	}
}
