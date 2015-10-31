package main

import "fmt"

// START OMIT
func Double(n int) int {
	return n * 2
}

func Call(fn func() int) int {
	return fn()
}

func main() {
	fmt.Println("Double 2:", Double(2))

	fmt.Println("Call Double 2000:", Call(func() int {
		return Double(2000)
	}))
}

// END OMIT
