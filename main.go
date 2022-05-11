package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Error: A CUE file is required")
		os.Exit(1)
	}
}
