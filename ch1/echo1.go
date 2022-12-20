package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}

	fmt.Println(s)
}
