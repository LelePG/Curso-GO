package main

import "fmt"

func main() {
	favSport:="chess"
	switch favSport {
	case "football":
		fmt.Println("You like football")

	case "baseball":
		fmt.Println("You like baseball")

	case "badminton":
		fmt.Println("You like badminton")

	case "tennis":
		fmt.Println("You like tennis")

	case "voley":
		fmt.Println("You like voley")

	case "racing":
		fmt.Println("You like racing")
	default:
		fmt.Println("I don't think that's a sport")
	}

}
