package main

import "fmt"

type square struct {
	lentgh int
	width int
}

type circle struct{
	r int
}

type shape interface{
	area() int
}


func (s square) area() int{
	a:=s.lentgh * s.width
	return a
}

func (c circle) area() int{
	a:=c.r*3*c.r
	return a
}

func  info(sh shape){
	fmt.Println(sh.area())
}

func main() {
	cc:=circle{r:5,}
	ss:=square{lentgh:1,width:2}
	info(cc)
	info(ss)

}