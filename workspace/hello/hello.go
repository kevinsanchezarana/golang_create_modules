package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))
	// You can even add new methods in third modules and use them: reverse.Int(24601)
}
