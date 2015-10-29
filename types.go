package main

import "fmt"

type Int int

func (i Int) Plus(j Int) Int {
	return i + j
}

func main() {
	i := Int(1)
	fmt.Println(i.Plus(Int(2)))
}
