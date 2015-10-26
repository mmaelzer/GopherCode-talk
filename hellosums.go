package main

import (
	"fmt"

	"github.com/mmaelzer/sums"
)

func main() {
	fmt.Println("Hello", sums.SumVariadic(1, 2, 3, 4, 5))
}
