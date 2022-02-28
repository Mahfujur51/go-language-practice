package main

import "fmt"

type Rect struct {
	len int
	wid int
}

func (r Rect) Area() int {
	return r.len * r.wid
}

func main() {
	r := Rect{10, 10}
	fmt.Println(r)
	fmt.Println(r.Area())
}
