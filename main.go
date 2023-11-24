package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout,  file)
	file.Close()
	
}