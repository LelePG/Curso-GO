package main

import "fmt"

type vehicle struct{
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct{
	vehicle
	luxury bool
}

func main() {
	t1 := truck {
		vehicle: vehicle{2, "white"},
		fourWheel:true,
	}

	t2 := sedan{
		vehicle: vehicle{2,"black"},
		luxury:  false,
	}
	fmt.Println(t1,t1.color)
	fmt.Println(t2,t2.color)
}
