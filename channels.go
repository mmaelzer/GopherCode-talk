package main

import "fmt"

// START OMIT
func fibonacci() chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i
		}
	}()
	return c
}

func main() {
	c := fibonacci()
	for n := 1; n <= 10; n++ {
		fval := <-c
		fmt.Printf("%d: %d\n", n, fval)
	}
}

// END OMIT
