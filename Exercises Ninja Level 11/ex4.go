package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	k, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(k)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err: errors.New( "Raiz de número negativo"),
		}
		return f, e

	}
	return 42, nil
}
