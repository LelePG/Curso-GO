package main
import "fmt"

func main() {
	year := 2000
	for {
		fmt.Println(year)
		year++
		if year > 2020{
			break
		}
	}
}

