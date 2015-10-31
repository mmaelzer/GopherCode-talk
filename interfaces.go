package main

import "fmt"

// START OMIT
type Firster interface {
	First() int
}

type IntList struct {
	List []int
}

func (i *IntList) First() int {
	return i.List[0]
}

func PrintFirst(l Firster) {
	fmt.Println("first:", l.First())
}

func main() {
	PrintFirst(&IntList{[]int{1, 2, 3, 4, 5}})
}

// END OMIT
