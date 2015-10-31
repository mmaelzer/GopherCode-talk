package main

import "fmt"

// START OMIT
type Dog struct {
	Name string
	Age  int
}

func (dog *Dog) Speak() {
	for i := 0; i < dog.Age; i++ {
		fmt.Println("bark")
	}
}

func main() {
	scoops := Dog{"scoops", 3}
	scoops.Speak()
}

// END OMIT
