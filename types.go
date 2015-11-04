package main

import "fmt"

type Int int

func (i Int) Plus(j Int) Int {
	return i + j
}

func main() {
	i := Int(1)
	j := Int(2)
	fmt.Println(i.Plus(j))
}
