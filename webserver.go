package main

// START OMIT
import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Println("=== hello server running on 127.0.0.1:8080 ===")
	http.ListenAndServe(":8080", nil)
}

// END OMIT
